package repositories

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/allergy"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/auth"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/contactInfo"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/doctor"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/emergencyReception"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/emergencyReceptionMedServices"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/medService"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/patient"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/patientsAllergy"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/personalInfo"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/reception"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	interfaces.AuthRepository
	interfaces.AllergyRepository
	interfaces.DoctorRepository
	interfaces.MedServiceRepository
	interfaces.EmergencyReceptionMedServicesRepository
	interfaces.PatientRepository
	interfaces.PatientsAllergyRepository
	interfaces.ContactInfoRepository
	interfaces.EmergencyReceptionRepository
	interfaces.PersonalInfoRepository
	interfaces.ReceptionRepository
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
	if err := db.AutoMigrate(
		&entities.Doctor{},
		&entities.Reception{},
		&entities.EmergencyReception{},
		// ... другие модели
	); err != nil {
		return nil, fmt.Errorf("ошибка выполнения автомиграций: %w", err)
	}

	// Создание тестовых данных
	if err := createTestDoctors(db); err != nil {
		return nil, fmt.Errorf("ошибка создания тестовых данных: %w", err)
	}

	return &Repository{
		auth.NewAuthRepository(db),
		allergy.NewAllergyRepository(db),
		doctor.NewDoctorRepository(db),
		medService.NewMedServiceRepository(db),
		emergencyReceptionMedServices.NewEmergencyReceptionMedServicesRepository(db),
		patient.NewPatientRepository(db),
		patientsAllergy.NewPatientsAllergyRepository(db),
		contactInfo.NewContactInfoRepository(db),
		emergencyReception.NewEmergencyReceptionRepository(db),
		personalInfo.NewPersonalInfoRepository(db),
		reception.NewReceptionRepository(db),
	}, nil
}

// createTestDoctors создает тестовых врачей при инициализации
func createTestDoctors(db *gorm.DB) error {
	testDoctors := []entities.Doctor{
		{
			FullName:       "Иванов Иван Иванович",
			Specialization: "Терапевт",
			Login:          "doctor1",
			PasswordHash:   hashPassword("password1"),
		},
		{
			FullName:       "Петров Петр Петрович",
			Specialization: "Хирург",
			Login:          "doctor2",
			PasswordHash:   hashPassword("password2"),
		},
		{
			FullName:       "Сидорова Анна Владимировна",
			Specialization: "Невролог",
			Login:          "doctor3",
			PasswordHash:   hashPassword("password3"),
		},
	}

	for _, doctor := range testDoctors {
		if err := db.FirstOrCreate(&doctor, entities.Doctor{Login: doctor.Login}).Error; err != nil {
			return err
		}
	}
	return nil
}

// hashPassword хэширует пароль для безопасного хранения
func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
