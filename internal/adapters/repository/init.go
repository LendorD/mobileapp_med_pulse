package repositories

import (
	"fmt"
	"log"
	"os"
	"time"

	emergencyReceptionMedServices "github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/EmergencyReceptionMedServices"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/allergy"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/contactInfo"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/doctor"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/emergencyReception"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/medService"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/patient"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/patientsAllergy"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/personalInfo"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repository/reception"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
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
		AllergyRepository:                       allergy.NewAllergyRepository(db),
		ContactInfoRepository:                   contactInfo.NewContactInfoRepository(db),
		DoctorRepository:                        doctor.NewDoctorRepository(db),
		EmergencyReceptionRepository:            emergencyReception.NewEmergencyReceptionRepository(db),
		EmergencyReceptionMedServicesRepository: emergencyReceptionMedServices.NewEmergencyReceptionMedServicesRepository(db),
		MedServiceRepository:                    medService.NewMedServiceRepository(db),
		PatientRepository:                       patient.NewPatientRepository(db),
		PatientsAllergyRepository:               patientsAllergy.NewPatientsAllergyRepository(db),
		PersonalInfoRepository:                  personalInfo.NewPersonalInfoRepository(db),
		ReceptionRepository:                     reception.NewReceptionRepository(db),
	}, nil

	return nil, nil
}

// autoMigrate - выполнение автомиграций для моделей
func autoMigrate(db *gorm.DB) error {
	models := []interface{}{
		&entities.Allergy{},
		&entities.Reception{},
		&entities.Patient{},
		&entities.ContactInfo{},
		&entities.PersonalInfo{},
		&entities.EmergencyReception{},
		&entities.EmergencyReceptionMedServices{},
		&entities.MedService{},
		&entities.PatientsAllergy{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("ошибка миграции модели %T: %w", model, err)
		}
	}
	/*
		// Добавляем начальные записи в таблицу BookingStatus
		initialStatuses := []entities.BookingStatus{
			{ID: 1, Name: "Создан"},
			{ID: 2, Name: "Выполняется"},
			{ID: 3, Name: "Выполняется без отклонений"},
			{ID: 4, Name: "Выполняется с отклонениями"},
			{ID: 5, Name: "Авария"},
			{ID: 6, Name: "Ошибка"},
			{ID: 7, Name: "Остановка"},
			{ID: 8, Name: "Выполнено"},
			{ID: 9, Name: "Отменено"},
		}

		for _, status := range initialStatuses {
			if err := db.FirstOrCreate(&status, entities.BookingStatus{ID: status.ID}).Error; err != nil {
				return fmt.Errorf("ошибка создания начальной записи статуса %v: %w", status, err)
			}
		}

			// Добавляем начальные записи в таблицу Specialty
			initialSpecialty := []entities.Specialty{
				{ID: 1, Name: "Металообработка"},
				{ID: 2, Name: "Гибридная обработка"},
			}

			for _, specialty := range initialSpecialty {
				if err := db.FirstOrCreate(&specialty, entities.Specialty{ID: specialty.ID}).Error; err != nil {
					return fmt.Errorf("ошибка создания начальной записи специализации %v: %w", specialty, err)
				}
			}
			// Добавляем начальные записи в таблицу GroupMaterial
			initialGroupMaterial := []entities.GroupMaterial{
				{ID: 1, Name: "Основная группа материалов"},
			}

			for _, groupMaterial := range initialGroupMaterial {
				if err := db.FirstOrCreate(&groupMaterial, entities.GroupMaterial{ID: groupMaterial.ID}).Error; err != nil {
					return fmt.Errorf("ошибка создания начальной записи группы материалов %v: %w", groupMaterial, err)
				}
			}
			// Добавляем начальные записи в таблицу GroupService
			initialGroupService := []entities.GroupService{
				{ID: 1, Name: "Основная группа номенклатуры"},
			}

			for _, groupService := range initialGroupService {
				if err := db.FirstOrCreate(&groupService, entities.GroupService{ID: groupService.ID}).Error; err != nil {
					return fmt.Errorf("ошибка создания начальной записи группы номенклатуры %v: %w", groupService, err)
				}
			}
			// Добавляем начальные записи в таблицу TypeService
			initialTypeService := []entities.TypeService{
				{ID: 1, Name: "Деталь"},
				{ID: 2, Name: "Заготовка"},
			}

			for _, typeService := range initialTypeService {
				if err := db.FirstOrCreate(&typeService, entities.TypeService{ID: typeService.ID}).Error; err != nil {
					return fmt.Errorf("ошибка создания начальной записи типов номенклатуры %v: %w", typeService, err)
				}
			}

			// Добавляем начальные записи в таблицу Unit
			initialUnit := []entities.Unit{
				{ID: 1, Name: "шт."},
				{ID: 2, Name: "кг."},
			}

			for _, unit := range initialUnit {
				if err := db.FirstOrCreate(&unit, entities.Unit{ID: unit.ID}).Error; err != nil {
					return fmt.Errorf("ошибка создания начальной записи единиц измерения %v: %w", unit, err)
				}
			}

	*/
	return nil
}
