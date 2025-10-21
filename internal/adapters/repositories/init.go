package repositories

import (
	"fmt"
	"log"
	"os"
	"time"

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
	}, nil

}

func autoMigrate(db *gorm.DB) error {
	// 🔥 Только для dev! Удаляем ВСЁ
	log.Println("🗑️ Dropping all tables...")

	// Сначала дочерние таблицы (с FK), потом родительские
	_ = db.Migrator().DropTable(&entities.OneCMedicalCard{})
	_ = db.Migrator().DropTable(&entities.OneCPatientListItem{})
	_ = db.Migrator().DropTable(&entities.OneCReception{})
	_ = db.Migrator().DropTable(&entities.AuthUser{})

	log.Println("🆕 Creating tables in correct order...")

	// Теперь создаём в правильном порядке
	if err := db.Migrator().CreateTable(&entities.AuthUser{}); err != nil {
		return fmt.Errorf("auth_users: %w", err)
	}
	if err := db.Migrator().CreateTable(&entities.OneCReception{}); err != nil {
		return fmt.Errorf("receptions: %w", err)
	}
	if err := db.Migrator().CreateTable(&entities.OneCPatientListItem{}); err != nil {
		return fmt.Errorf("patient_list: %w", err)
	}
	if err := db.Migrator().CreateTable(&entities.OneCMedicalCard{}); err != nil {
		return fmt.Errorf("med_cards: %w", err)
	}

	log.Println("✅ Migrations completed")
	return nil
}

func seedInitialData(db *gorm.DB, cfg *config.Config) error {
	// Проверяем, есть ли уже демо-данные (например, по первому пользователю)
	var count int64
	db.Model(&entities.AuthUser{}).Where("login = ?", "user1").Count(&count)
	if count > 0 {
		log.Println("ℹ️ Demo data already exists, skipping seeding")
		return nil
	}

	log.Println("🌱 Seeding initial demo data...")

	// Хешируем пароль один раз (для всех пользователей — одинаковый)
	password := "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Собираем данные для пакетной вставки
	var authUsers []entities.AuthUser
	var medicalCards []entities.OneCMedicalCard
	var patientListItems []entities.OneCPatientListItem

	for i := 1; i <= 10; i++ {
		login := fmt.Sprintf("+7962284076%d", i)
		patientID := fmt.Sprintf("user%d_id", i)
		fullName := fmt.Sprintf("Пациент %d", i)

		// 1. Пользователь аутентификации
		authUsers = append(authUsers, entities.AuthUser{
			Login:    login,
			Password: string(hash),
		})

		// 2. Медицинская карта
		medicalCards = append(medicalCards, entities.OneCMedicalCard{
			PatientID:   patientID,
			DisplayName: fullName,
			Age:         fmt.Sprintf("%d", 20+i%50),
			BirthDate:   fmt.Sprintf("198%d-0%d-1%d", i%9+1, i%12+1, i%28+1),
			MobilePhone: fmt.Sprintf("+790012345%02d", i),
			Address:     fmt.Sprintf("г. Москва, ул. Тестовая, д. %d", i),
			Email:       fmt.Sprintf("user%d@example.com", i),
			Workplace:   fmt.Sprintf("ООО \"Компания %d\"", i),
			Snils:       fmt.Sprintf("123-456-789 %d", i),

			LegalRepresentative: entities.ClientRef{
				ID:   fmt.Sprintf("rep_%d", i),
				Name: fmt.Sprintf("Представитель %d", i),
			},
			Relative: entities.Relative{
				Status: "Родственник",
				Name:   fmt.Sprintf("Родственник %d", i),
			},
			AttendingDoctor: entities.Doctor{
				FullName:           fmt.Sprintf("Доктор %d", i),
				PolicyOrCertNumber: fmt.Sprintf("POL%d", i),
				AttachmentStart:    "2020-01-01",
				AttachmentEnd:      "2030-01-01",
				Clinic:             fmt.Sprintf("Поликлиника %d", i),
			},
			Policy: entities.Policy{
				Number: fmt.Sprintf("POLICY%d", i),
				Type:   "ОМС",
			},
			Certificate: entities.Certificate{
				Number: fmt.Sprintf("CERT%d", i),
				Date:   "2023-01-01",
			},
		})

		// 3. Элемент списка пациентов
		patientListItems = append(patientListItems, entities.OneCPatientListItem{
			PatientID: patientID,
			FullName:  fullName,
			Gender:    i%2 == 0, // чередуем пол
			BirthDate: fmt.Sprintf("198%d-0%d-1%d", i%9+1, i%12+1, i%28+1),
		})
	}

	// 1. Пользователи аутентификации
	if err := db.CreateInBatches(authUsers, 10).Error; err != nil {
		return fmt.Errorf("failed to seed auth users: %w", err)
	}

	// 2. Список пациентов (родительская таблица для медкарт!)
	if err := db.CreateInBatches(patientListItems, 10).Error; err != nil {
		return fmt.Errorf("failed to seed patient list items: %w", err)
	}

	// 3. Медицинские карты (дочерняя таблица)
	if err := db.CreateInBatches(medicalCards, 10).Error; err != nil {
		return fmt.Errorf("failed to seed medical cards: %w", err)
	}

	// Добавим одну заявку на скорую (опционально)
	emergencyCall := entities.OneCReception{
		CallID: "demo_call_001",
		Status: "received",
		Data:   []byte(`{"patient": "demo", "reason": "test"}`),
	}
	if err := db.Create(&emergencyCall).Error; err != nil {
		log.Printf("⚠️ Warning: failed to seed emergency call: %v", err)
	}

	log.Println("✅ Demo data seeded successfully (10 users, 10 medical cards, 10 patient list items)")
	return nil
}
