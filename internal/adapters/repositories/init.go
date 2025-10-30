package repositories

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/analysis"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/file"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/flg"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/medcard"
	receptionSmp "github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/reception_smp"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/tx"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/auth"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/doctor"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/patient"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	interfaces.AuthRepository
	interfaces.DoctorRepository
	interfaces.PatientRepository
	interfaces.ReceptionSmpRepository
	interfaces.MedicalCardRepository
	interfaces.TxManager
	interfaces.FileRepository
	interfaces.FlgRepository
	interfaces.AnalysisRepository
}

func NewRepository(cfg *config.Config) (interfaces.Repository, error) {
	//logger := logging.NewModuleLogger("ADAPTER", "POSTGRES", parentLogger)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Вывод в stdout
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Порог для медленных запросов
			LogLevel:                  logger.Info,            // Уровень логирования (Info - все запросы)
			IgnoreRecordNotFoundError: true,                   // Игнорировать ошибки "запись не найдена"
			Colorful:                  true,                   // Цветной вывод
		},
	)

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Выполнение автомиграций
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("ошибка выполнения автомиграций: %w", err)
	}

	if err := seedInitialData(db, cfg); err != nil {
		log.Printf("⚠️ Ошибка сидов: %v", err)
		// Не падаем — сиды не критичны
	}

	return &Repository{
		auth.NewAuthRepository(db),
		doctor.NewDoctorRepository(db),
		patient.NewPatientRepository(db),
		receptionSmp.NewReceptionSmpRepository(db),
		medcard.NewMedicalCardRepository(db),
		tx.NewTxManager(db),
		file.NewFileRepository(db),
		flg.NewFlgRepository(db),
		analysis.NewAnalysisRepository(db),
	}, nil

}

func autoMigrate(db *gorm.DB) error {
	log.Println("🗑️ Dropping all tables (dev mode)...")

	// Удаляем ВСЁ в обратном порядке (сначала дочерние, потом родительские)
	tables := []interface{}{
		&entities.PatientData{},
		&entities.DoctorData{},
		&entities.Receptions{},
		&entities.Flg{},
		&entities.Analysis{},
		&entities.OneCReception{},
		&entities.OneCMedicalCard{},
		&entities.OneCPatientListItem{},
		&entities.File{},
		&entities.AuthUser{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Printf("⚠️ Ошибка при удалении таблицы: %v", err)
		}
	}

	log.Println("🆕 Creating tables in correct order...")

	// 1. Пользователи системы (авторизация)
	if err := db.AutoMigrate(&entities.AuthUser{}); err != nil {
		return fmt.Errorf("auth_users: %w", err)
	}

	// 2. Файлы (используются в Flg и других местах)
	if err := db.AutoMigrate(&entities.File{}); err != nil {
		return fmt.Errorf("files: %w", err)
	}

	// 3. Основной вызов
	if err := db.AutoMigrate(&entities.OneCReception{}); err != nil {
		return fmt.Errorf("one_c_receptions: %w", err)
	}

	// 4. Дочерние сущности вызова (ссылаются на OneCReception)
	if err := db.AutoMigrate(
		&entities.PatientData{},
		&entities.DoctorData{},
		&entities.Receptions{},
		&entities.Flg{},
		&entities.Analysis{},
	); err != nil {
		return fmt.Errorf("child tables of reception: %w", err)
	}

	// 5. Справочник пациентов (независимый)
	if err := db.AutoMigrate(&entities.OneCPatientListItem{}); err != nil {
		return fmt.Errorf("one_c_patient_list_items: %w", err)
	}

	// 6. Медицинские карты (ссылаются на OneCPatientListItem по PatientID)
	if err := db.AutoMigrate(&entities.OneCMedicalCard{}); err != nil {
		return fmt.Errorf("one_c_medical_cards: %w", err)
	}

	log.Println("✅ All tables migrated successfully")
	return nil
}

func seedInitialData(db *gorm.DB, cfg *config.Config) error {
	// Проверяем, есть ли уже демо-данные
	var authCount int64
	db.Model(&entities.AuthUser{}).Count(&authCount)
	if authCount > 0 {
		log.Println("ℹ️ Demo data already exists, skipping seeding")
		return nil
	}

	log.Println("🌱 Seeding initial demo data...")

	// === 1. Хешируем пароль ===
	password1 := "123"
	hash1, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	password2 := "321"
	hash2, err := bcrypt.GenerateFromPassword([]byte(password2), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	// === 2. Создаём пользователей системы (врачей) ===
	authUsers := []entities.AuthUser{
		{Login: "+79622840765", Password: string(hash1)}, // Врач 1
		{Login: "+79622840766", Password: string(hash2)}, // Врач 2
	}
	if err := db.CreateInBatches(authUsers, len(authUsers)).Error; err != nil {
		return fmt.Errorf("failed to seed auth users: %w", err)
	}
	log.Println("✅ Seeded 2 auth users")

	// === 3. Создаём пациентов (справочник) ===
	var patientListItems []entities.OneCPatientListItem
	var medicalCards []entities.OneCMedicalCard

	for i := 1; i <= 5; i++ {
		patientID := fmt.Sprintf("PATIENT_%03d", i)
		fullName := fmt.Sprintf("Пациент Иванович %d", i)

		// Справочник пациентов
		patientListItems = append(patientListItems, entities.OneCPatientListItem{
			PatientID: patientID,
			FullName:  fullName,
			Gender:    i%2 == 0,
			BirthDate: fmt.Sprintf("198%d-0%d-1%d", i%9+1, i%12+1, i%28+1),
		})

		// Медицинская карта
		medicalCards = append(medicalCards, entities.OneCMedicalCard{
			PatientID:   patientID,
			DisplayName: fullName,
			Age:         fmt.Sprintf("%d", 25+i%40),
			BirthDate:   fmt.Sprintf("198%d-0%d-1%d", i%9+1, i%12+1, i%28+1),
			MobilePhone: fmt.Sprintf("+790012345%02d", i),
			Address:     fmt.Sprintf("г. Москва, ул. Лечебная, д. %d", i),
			Email:       fmt.Sprintf("patient%d@example.com", i),
			Workplace:   fmt.Sprintf("ООО \"Здоровье %d\"", i),
			Snils:       fmt.Sprintf("123-456-789 %d", i),

			LegalRepresentative: entities.ClientRef{
				ID:   fmt.Sprintf("REP_%d", i),
				Name: fmt.Sprintf("Представитель %d", i),
			},
			Relative: entities.Relative{
				Status: "Мать",
				Name:   fmt.Sprintf("Мама Пациента %d", i),
			},
			Policy: entities.Policy{
				Number: fmt.Sprintf("POLICY_OMS_%d", i),
				Type:   "ОМС",
			},
			Certificate: entities.Certificate{
				Number: fmt.Sprintf("CERT_%d", i),
				Date:   "2023-05-15",
			},
		})
	}

	// Сохраняем пациентов
	if err := db.CreateInBatches(patientListItems, len(patientListItems)).Error; err != nil {
		return fmt.Errorf("failed to seed patient list: %w", err)
	}
	log.Println("✅ Seeded 5 patients")

	// Сохраняем медкарты
	if err := db.CreateInBatches(medicalCards, len(medicalCards)).Error; err != nil {
		return fmt.Errorf("failed to seed medical cards: %w", err)
	}
	log.Println("✅ Seeded 5 medical cards")

	// === 4. (Опционально) Создаём пример вызова скорой ===
	reception := entities.OneCReception{
		CallID:  "CALL_001",
		Address: "г. Москва, ул. Скорой помощи, д. 1",
		Phone:   "+79001112233",
		Status:  "received",
	}

	// Создаём вызов
	if err := db.Create(&reception).Error; err != nil {
		return fmt.Errorf("failed to seed reception: %w", err)
	}

	// Дочерние сущности
	patients := []entities.PatientData{
		{
			ReceptionID: reception.ID,
			PatientID:   "PATIENT_001",
			FullName:    "Пациент Иванович 1",
			Age:         45,
			BirthDate:   "1979-05-12",
			MobilePhone: "+79001234501",
			Policy: entities.Policy{
				Number: "POLICY_OMS_001",
				Type:   "ОМС",
			},
			Certificate: entities.Certificate{
				Number: "CERT_001",
				Date:   "2023-01-10",
			},
		},
	}

	doctor := entities.DoctorData{
		ReceptionID:    reception.ID,
		Name:           "Докторова Анна Петровна",
		Specialization: "Терапевт",
	}

	flg := entities.Flg{
		ReceptionID:  reception.ID,
		PatientID:    1,
		Organization: "Городская больница №1",
		Number:       "FLG-2025-001",
		Result:       "Без патологий",
		Date:         time.Now(),
	}

	analysis := entities.Analysis{
		ReceptionID: reception.ID,
		Code:        "A001",
		Title:       "Общий анализ крови",
		Price:       500,
	}

	// receptions := []entities.Receptions{
	// 	{
	// 		ReceptionID: reception.ID,
	// 		Data:        []byte(`{"diagnosis": "ОРВИ", "recommendations": "Постельный режим"}`),
	// 	},
	// }

	// Сохраняем всё
	if err := db.CreateInBatches(patients, len(patients)).Error; err != nil {
		return fmt.Errorf("failed to seed reception patients: %w", err)
	}
	if err := db.Create(&doctor).Error; err != nil {
		return fmt.Errorf("failed to seed doctor: %w", err)
	}
	if err := db.Create(&flg).Error; err != nil {
		return fmt.Errorf("failed to seed flg: %w", err)
	}
	if err := db.Create(&analysis).Error; err != nil {
		return fmt.Errorf("failed to seed analysis: %w", err)
	}
	// if err := db.CreateInBatches(receptions, len(receptions)).Error; err != nil {
	// 	return fmt.Errorf("failed to seed receptions: %w", err)
	// }

	log.Println("✅ Seeded 1 emergency call with all related data")

	log.Println("🎉 Demo data seeded successfully!")
	return nil
}
