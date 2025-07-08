package models

<<<<<<< HEAD
import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID             uint      `json:"id"`
	FirstName      string    `json:"first_name"`
	MiddleName     string    `json:"middle_name"`
	LastName       string    `json:"last_name"`
	Login          string    `json:"login"`
	PasswordHash   string    `json:"password_hash"`
	Specialization string    `json:"specialization"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Patient struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	BirthDate time.Time `json:"birth_date"` // Лучше использовать time.Time вместо int
	IsMale    bool      `json:"is_male"`    // true - мужской, false - женский
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContactInfo struct {
	ID        uint      `json:"id"`
	PatientID uint      `json:"patient_id"`
	Phone     string    `json:"phone"` // Строка лучше для телефонных номеров
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReceptionStatus string

const (
	StatusScheduled ReceptionStatus = "scheduled" // Запланирован
	StatusCompleted ReceptionStatus = "completed" // Завершен
	StatusCancelled ReceptionStatus = "cancelled" // Отменен
	StatusNoShow    ReceptionStatus = "no_show"   // Не явился
)

type Reception struct {
	ID              uint            `json:"id"`
	DoctorID        uint            `json:"doctor_id"`
	PatientID       uint            `json:"patient_id"`
	Date            time.Time       `json:"date"`
	Diagnosis       string          `json:"diagnosis"`       // Диагноз
	Recommendations string          `json:"recommendations"` // Рекомендации
	Status          ReceptionStatus `json:"status"`
	Address         string          `json:"address"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type ReceptionSmpStatus string

const (
	StatusEmergency ReceptionSmpStatus = "emergency" // 1 Приоритеты
	StatusPriority  ReceptionSmpStatus = "priority"  // 2
	StatusUrgent    ReceptionSmpStatus = "urgent"    // 3
)

type ReceptionSMP struct {
	ID          uint               `json:"id"`
	ReceptionID uint               `json:"doctor_id"`
	Status      ReceptionSmpStatus `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type Allergy struct {
	ID          uint      `json:"id"`
	PatientID   uint      `json:"patient_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RecordedAt  time.Time `json:"recorded_at"`
}

type PatientCard struct {
	Patient     Patient     `json:"patient"`
	ContactInfo ContactInfo `json:"contact_info"`
	Receptions  []Reception `json:"receptions"` // Массив (слайс) структур Receptions
	Allergies   []Allergy   `json:"allergies"`  // Массив структур аллергий
}

// Для JWT-авторизации
type User struct {
	gorm.Model
	Login        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"unique"`
	Role         string `gorm:"default:'user'"`
}

type Session struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	RefreshToken string `gorm:"not null"`
	ExpiresAt    int64  `gorm:"not null"`
=======
import "time"

type EmergencyStatus string

const (
	EmergencyStatusCritical EmergencyStatus = "emergency" // "Экстренный"
	EmergencyStatusUrgent   EmergencyStatus = "urgent"    // "Неотложный"
)

type EmergencyPatient struct {
	FullName    string          `json:"full_name" example:"Смирнов Александр Петрович" rus:"ФИО пациента"`
	Address     string          `json:"address" example:"ул. Ленина, д. 15, кв. 42" rus:"Адрес"`
	Doctor      string          `json:"doctor" example:"Иванова М.П." rus:"Лечащий врач"`
	Status      EmergencyStatus `json:"status" example:"emergency" rus:"Статус"`
	ArrivalTime time.Time       `json:"arrival_time" example:"2023-10-15T14:30:00Z" rus:"Время поступления"`
}

type EmergencyListResponse struct {
	Patients []EmergencyPatient `json:"patients" rus:"Список экстренных пациентов"`
}

type PatientShort struct {
	ID            uint   `json:"id" example:"1" rus:"ID пациента"`
	FullName      string `json:"full_name" example:"Иванов Иван Иванович" rus:"ФИО"`
	RoomNumber    string `json:"room_number" example:"101" rus:"Номер палаты"`
	MainDiagnosis string `json:"main_diagnosis" example:"Гипертоническая болезнь II ст." rus:"Основной диагноз"`
}

type PatientListResponse struct {
	Patients []PatientShort `json:"patients" rus:"Список пациентов"`
}

type MedicalRecord struct {
	Date            string `json:"date" example:"12.05.2023" rus:"Дата приема"`
	Diagnosis       string `json:"diagnosis" example:"ОРВИ, острая форма" rus:"Диагноз"`
	Recommendations string `json:"recommendations" example:"Постельный режим, обильное питье" rus:"Рекомендации"`
	Doctor          string `json:"doctor" example:"Терапевт Иванова А.П." rus:"Врач"`
}

type MedicalHistoryResponse struct {
	Records []MedicalRecord `json:"records" rus:"История болезни"`
}

type PatientFull struct {
	FullName          string `json:"full_name" example:"Иванов Иван Иванович" rus:"ФИО"`
	Gender            string `json:"gender" example:"Мужской" rus:"Пол"`
	BirthDate         string `json:"birth_date" example:"15.03.1965" rus:"Дата рождения"`
	SNILS             string `json:"snils" example:"123-456-789 01" rus:"СНИЛС"`
	OMS               string `json:"oms" example:"1234567890123456" rus:"Полис ОМС"`
	Passport          string `json:"passport" example:"4510 123456" rus:"Паспорт"`
	Phone             string `json:"phone" example:"+7 (999) 123-45-67" rus:"Телефон"`
	Email             string `json:"email" example:"ivanov@example.com" rus:"Email"`
	Address           string `json:"address" example:"г. Москва, ул. Ленина, д. 15, кв. 42" rus:"Адрес"`
	Contraindications string `json:"contraindications" example:"Аллергия на пенициллин" rus:"Противопоказания"`
}

type PatientCreateRequest struct {
	FullName          string `json:"full_name" binding:"required" example:"Иванов Иван Иванович" rus:"ФИО"`
	Gender            string `json:"gender" binding:"required" example:"Мужской" rus:"Пол"`
	BirthDate         string `json:"birth_date" binding:"required" example:"15.03.1965" rus:"Дата рождения"`
	Passport          string `json:"passport" binding:"required" example:"4510 123456" rus:"Паспорт"`
	SNILS             string `json:"snils" binding:"required" example:"123-456-789 01" rus:"СНИЛС"`
	OMS               string `json:"oms" binding:"required" example:"1234567890123456" rus:"Полис ОМС"`
	Phone             string `json:"phone" binding:"required" example:"+7 (999) 123-45-67" rus:"Телефон"`
	Email             string `json:"email" example:"ivanov@example.com" rus:"Email"`
	Address           string `json:"address" binding:"required" example:"г. Москва, ул. Ленина, д. 15, кв. 42" rus:"Адрес"`
	Contraindications string `json:"contraindications" example:"Аллергия на пенициллин" rus:"Противопоказания"`
}

type DoctorResponse struct {
	ID             uint   `json:"id" example:"1" rus:"ID врача"`
	FullName       string `json:"full_name" example:"Иванов Иван Иванович" rus:"ФИО"`
	Specialization string `json:"specialization" example:"Терапевт" rus:"Специализация"`
}

type DoctorRegisterRequest struct {
	FullName        string `json:"full_name" binding:"required" example:"Иванов Иван Иванович" rus:"ФИО"`
	Login           string `json:"login" binding:"required" example:"doctor_ivanov" rus:"Логин"`
	Password        string `json:"password" binding:"required,min=8" example:"securepassword123" rus:"Пароль"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password" example:"securepassword123" rus:"Подтверждение пароля"`
	Specialization  string `json:"specialization" binding:"required" example:"Терапевт" rus:"Специализация"`
}

type Diagnosis struct {
	Name        string `json:"name" example:"Гипертоническая болезнь II ст." rus:"Название"`
	Description string `json:"description" example:"Умеренная артериальная гипертензия" rus:"Описание"`
}

type RoomInfo struct {
	Number string `json:"number" example:"101" rus:"Номер палаты"`
	Type   string `json:"type" example:"2-местная" rus:"Тип палаты"`
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"404" rus:"Код ошибки"`
	Message string `json:"message" example:"Not found" rus:"Сообщение"`
	Details string `json:"details,omitempty" example:"patient not found" rus:"Детали"`
>>>>>>> main
}
