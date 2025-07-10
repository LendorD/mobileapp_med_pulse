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
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"

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
	/*
		// Выполнение автомиграций
		if err := autoMigrate(db); err != nil {
			return nil, fmt.Errorf("ошибка выполнения автомиграций: %w", err)

		}
	*/
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

	//return nil, nil
}
