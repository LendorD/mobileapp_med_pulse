package repositories

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/allergy"
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

	// Выполнение автомиграций
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("ошибка выполнения автомиграций: %w", err)

	}

	return &Repository{
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

// autoMigrate - выполнение автомиграций для моделей
func autoMigrate(db *gorm.DB) error {
	models := []interface{}{
		&entities.Doctor{},
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

	if err := dropTables(db); err != nil {
		return fmt.Errorf("ошибка удаления таблиц: %w", err)
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("ошибка миграции модели %T: %w", model, err)
		}
	}

	// Правильный вызов функции seedTestData
	if err := seedTestData(db); err != nil {
		return fmt.Errorf("ошибка заполнения тестовыми данными: %w", err)
	}

	return nil
}

func dropTables(db *gorm.DB) error {
	tables := []string{
		"receptions",
		"contact_infos",
		"personal_infos",
		"patients",
		"doctors",
	}

	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", strings.Join(tables, ", "))

	if err := db.Exec(query).Error; err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	return nil
}

func seedTestData(db *gorm.DB) error {
	// 1. Создаем докторов
	doctors := []*entities.Doctor{
		{
			FullName:       "Иванов Иван Иванович",
			Login:          "doctor_ivanov",
			Email:          "ivanov@clinic.ru",
			PasswordHash:   "$2a$10$somehashedpassword", // Пример хэша
			Specialization: "Терапевт",
		},
		{
			FullName:       "Петрова Мария Сергеевна",
			Login:          "doctor_petrova",
			Email:          "petrova@clinic.ru",
			PasswordHash:   "$2a$10$somehashedpassword",
			Specialization: "Хирург",
		},
		{
			FullName:       "Сидоров Алексей Дмитриевич",
			Login:          "doctor_sidorov",
			Email:          "sidorov@clinic.ru",
			PasswordHash:   "$2a$10$somehashedpassword",
			Specialization: "Кардиолог",
		},
	}

	for _, doc := range doctors {
		if err := db.Create(doc).Error; err != nil {
			continue
		}
	}

	// 2. Создаем пациентов
	patients := []*entities.Patient{
		{FullName: "Смирнов Алексей Петрович", BirthDate: parseDate("1980-05-15"), IsMale: true},
		{FullName: "Кузнецова Анна Владимировна", BirthDate: parseDate("1992-08-21"), IsMale: false},
		{FullName: "Попов Дмитрий Игоревич", BirthDate: parseDate("1975-11-03"), IsMale: true},
		{FullName: "Васильева Елена Александровна", BirthDate: parseDate("1988-07-14"), IsMale: false},
		{FullName: "Новиков Сергей Олегович", BirthDate: parseDate("1995-02-28"), IsMale: true},
		{FullName: "Морозова Ольга Дмитриевна", BirthDate: parseDate("1983-09-17"), IsMale: false},
		{FullName: "Лебедев Андрей Николаевич", BirthDate: parseDate("1978-12-05"), IsMale: true},
		{FullName: "Соколова Татьяна Викторовна", BirthDate: parseDate("1990-04-30"), IsMale: false},
		{FullName: "Козлов Артем Сергеевич", BirthDate: parseDate("1987-06-22"), IsMale: true},
		{FullName: "Павлова Наталья Игоревна", BirthDate: parseDate("1993-03-11"), IsMale: false},
	}

	for _, pat := range patients {
		if err := db.Create(pat).Error; err != nil {
			continue
		}
	}

	// 3. Создаем контактную информацию и персональные данные для пациентов
	for i, patient := range patients {
		contactInfo := entities.ContactInfo{
			PatientID: patient.ID, // теперь тут правильный ID
			Phone:     fmt.Sprintf("+7915%07d", 1000000+i),
			Email:     fmt.Sprintf("patient%d@example.com", i+1),
			Address:   fmt.Sprintf("Москва, ул. Тестовая, д. %d", i+1),
		}
		personalInfo := entities.PersonalInfo{
			PatientID:      patient.ID,
			PassportSeries: fmt.Sprintf("4510 %06d", 100000+i),
			SNILS:          fmt.Sprintf("123-456-789 %02d", i),
			OMS:            fmt.Sprintf("1234567890%d", i),
		}

		if err := db.Create(&contactInfo).Error; err != nil {
			return err
		}
		if err := db.Create(&personalInfo).Error; err != nil {
			return err
		}

		// Обновляем пациента с ID контактной информации
		db.Model(&patient).Updates(map[string]interface{}{
			"ContactInfoID":  &contactInfo.ID,
			"PersonalInfoID": &personalInfo.ID,
		})
	}

	// 4. Создаем приемы на 10, 11 и 12 июля текущего года
	now := time.Now()
	dates := []time.Time{
		time.Date(now.Year(), 7, 10, 0, 0, 0, 0, time.UTC),
		time.Date(now.Year(), 7, 11, 0, 0, 0, 0, time.UTC),
		time.Date(now.Year(), 7, 12, 0, 0, 0, 0, time.UTC),
	}

	statuses := []entities.ReceptionStatus{entities.StatusScheduled, entities.StatusCompleted, entities.StatusCancelled, entities.StatusNoShow}
	addresses := []string{
		"Москва, ул. Ленина, д. 15",
		"Москва, ул. Пушкина, д. 10",
		"Москва, пр. Вернадского, д. 25",
	}

	for i := 0; i < 50; i++ {
		// Выбираем случайные данные
		date := dates[i%len(dates)]
		hour := 9 + i%8 // Время приема с 9:00 до 16:00
		date = date.Add(time.Hour * time.Duration(hour))

		reception := entities.Reception{
			DoctorID:        doctors[i%len(doctors)].ID,
			PatientID:       patients[i%len(patients)].ID,
			Date:            date,
			Diagnosis:       "ОРВИ",
			Recommendations: "Постельный режим",
			IsOut:           i%3 == 0, // Каждый третий - выездной
			Status:          statuses[i%len(statuses)],
			Address:         addresses[i%len(addresses)],
		}

		if err := db.Create(&reception).Error; err != nil {
			return err
		}
	}

	statusesE := []entities.EmergencyStatus{
		entities.EmergencyStatusScheduled,
		entities.EmergencyStatusAccepted,
		entities.EmergencyStatusOnPlace,
		entities.EmergencyStatusCompleted,
		entities.EmergencyStatusCancelled,
		entities.EmergencyStatusNoShow,
	}

	for i := 0; i < 50; i++ {
		// Выбираем случайные данные
		date := dates[i%len(dates)]
		hour := 9 + i%8        // Время приема с 9:00 до 16:00
		minute := 30 * (i % 2) // 0 или 30 минут
		date = date.Add(time.Hour * time.Duration(hour)).Add(time.Minute * time.Duration(minute))

		emergencyReception := entities.EmergencyReception{
			DoctorID:  doctors[i%len(doctors)].ID,
			PatientID: patients[i%len(patients)].ID,
			Date:      date,
			Status:    statusesE[i%len(statusesE)],
			Priority:  i%2 == 0, // Каждый второй - экстренный (true), остальные - неотложные (false)
			Address:   addresses[i%len(addresses)],
		}

		if err := db.Create(&emergencyReception).Error; err != nil {
			return err
		}
	}

	return nil
}

func parseDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}
