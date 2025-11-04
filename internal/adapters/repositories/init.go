package repositories

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/analysis"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/emk"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/file"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/flg"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/medcard"
	receptionSmp "github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/reception_smp"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/tx"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

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
	interfaces.EmkRepository
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
	if err := Migrate(db); err != nil {
		return nil, fmt.Errorf("ошибка выполнения автомиграций: %w", err)
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
		emk.NewEmkRepository(db),
	}, nil

}

// Migrate выполняет миграции и заполняет тестовыми данными.
func Migrate(db *gorm.DB) error {
	// Автоматическая миграция
	err := db.AutoMigrate(
		&entities.AuthUser{},
		&entities.Emk{},
		&entities.Flg{},
		&entities.File{},
		&entities.OneCMedicalCard{},
		&entities.OneCPatientListItem{},
	)
	if err != nil {
		return err
	}

	log.Println("✅ Таблицы успешно смигрированы")

	// Проверим, есть ли уже данные
	var count int64
	db.Model(&entities.OneCPatientListItem{}).Count(&count)
	if count > 0 {
		log.Println("ℹ️  Тестовые данные уже существуют, пропуск заполнения")
		return nil
	}

	log.Println("🚀 Добавляем тестовые данные...")

	// --- Очистка данных перед перезапуском ---
	log.Println("🧹 Очищаем таблицы перед заполнением...")
	if err := truncateAll(db); err != nil {
		return err
	}

	// --- 1. Создаём 2 врача
	doctors := []entities.DoctorData{
		{ReceptionID: 1, Name: "Иванов Сергей Петрович", Specialization: "Терапевт"},
		{ReceptionID: 2, Name: "Петров Алексей Николаевич", Specialization: "Хирург"},
	}

	// --- 2. 10 пациентов с медкартами
	for i := 1; i <= 10; i++ {
		patient := entities.OneCPatientListItem{
			PatientID: generatePatientID(i),
			FullName:  generateName(i),
			Gender:    i%2 == 0,
			BirthDate: "1990-01-01",
		}
		db.Create(&patient)

		// Создаём медкарту
		card := entities.OneCMedicalCard{
			PatientID:       patient.PatientID,
			DisplayName:     patient.FullName,
			Age:             "35",
			BirthDate:       patient.BirthDate,
			MobilePhone:     "+79991234567",
			AdditionalPhone: "+79990001122",
			Address:         "г. Москва, ул. Пушкина, д. Колотушкина",
			Email:           "patient" + patient.PatientID + "@mail.ru",
			Workplace:       "ООО «Пример»",
			Snils:           "123-456-789 00",
			Doctor: entities.AttendingDoctor{
				FullName:           doctors[i%2].Name,
				Specialization:     doctors[i%2].Specialization,
				PolicyOrCertNumber: "ABC-" + patient.PatientID,
				AttachmentStart:    "2020-01-01",
				AttachmentEnd:      "2030-01-01",
				Clinic:             "Клиника №" + generateClinic(i),
			},
			Policy: entities.Policy{
				Number: "P-" + patient.PatientID,
				Type:   "ОМС",
			},
			Certificate: entities.Certificate{
				Number: "C-" + patient.PatientID,
				Date:   "2022-01-01",
			},
		}
		db.Create(&card)

		// Создаём вызов Emk
		emk := entities.Emk{
			PatientID:   patient.PatientID,
			CallID:      "CALL-" + patient.PatientID,
			Status:      entities.CallStatusWork,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			MedServices: []byte(`["Терапия", "Диагностика"]`),
		}
		db.Create(&emk)
	}

	log.Println("✅ Тестовые данные успешно добавлены!")
	return nil
}

// ---- Утилиты для тестов ----

// truncateAll удаляет все записи из таблиц.
func truncateAll(db *gorm.DB) error {
	tables := []string{
		"auth_users",
		"files",
		"flgs",
		"emks",
		"one_c_medical_cards",
		"one_c_patient_list_items",
	}
	for _, t := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", t)).Error; err != nil {
			return fmt.Errorf("ошибка очистки таблицы %s: %v", t, err)
		}
	}
	return nil
}

func generatePatientID(i int) string {
	return "P-" + fmt.Sprint(1000+i)
}

func generateName(i int) string {
	names := []string{
		"Иванов Иван", "Петров Петр", "Сидоров Алексей", "Кузнецов Дмитрий",
		"Смирнов Николай", "Федоров Сергей", "Егоров Михаил", "Ковалев Андрей",
		"Новиков Артем", "Лебедев Павел",
	}
	if i <= len(names) {
		return names[i-1]
	}
	return fmt.Sprintf("Пациент %d", i)
}

func generateClinic(i int) string {
	return fmt.Sprint(i%3 + 1)
}
