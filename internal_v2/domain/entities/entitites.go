package entities

import (
	"time"

	"gorm.io/gorm"
)

//type User struct {
//	gorm.Model
//	Login        string `gorm:"unique;not null" json:"login" example:"doctor_ivanov" rus:"Логин"`
//	PasswordHash string `gorm:"not null" json:"-" rus:"Хэш пароля"`
//
//}

/*
type Session struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id" example:"1" rus:"ID пользователя"`
	RefreshToken string `gorm:"not null" json:"refresh_token" example:"eyJhbGci..." rus:"Refresh токен"`
	ExpiresAt    int64  `gorm:"not null" json:"expires_at" example:"1735689600" rus:"Время истечения"`
}

*/

type Doctor struct {
	gorm.Model
	FullName       string `gorm:"not null" json:"-" example:"Иванов Иван Иванович" rus:"ФИО"`
	Login          string `gorm:"unique;not null" json:"login" example:"doctor_ivanov" rus:"Логин"`
	Email          string `gorm:"unique" json:"email" example:"ivanov@clinic.ru" rus:"Email"`
	PasswordHash   string `gorm:"not null" json:"-" rus:"Хэш пароля"`
	Specialization string `gorm:"not null" json:"specialization" example:"Терапевт" rus:"Специализация"`

	Reception          []Reception          `gorm:"foreignKey:ReceptionID"`
	EmergencyReception []EmergencyReception `gorm:"foreignKey:EmergencyReceptionID"`
}

type Patient struct {
	gorm.Model
	FullName  string    `gorm:"not null" json:"-" example:"Смирнов Алексей Петрович" rus:"ФИО"`
	BirthDate time.Time `gorm:"not null" json:"birth_date" example:"1980-05-15T00:00:00Z" rus:"Дата рождения"`
	IsMale    bool      `gorm:"not null" json:"is_male" example:"true" rus:"Пол (true - мужской)"`

	PersonalInfo PersonalInfo `gorm:"foreignKey:PersonalInfoID" json:"personal_info" rus:"Персональные данные"`
	ContactInfo  ContactInfo  `gorm:"foreignKey:ContactInfoID" json:"contact_info" rus:"Контактные данные"`

	Reception          []Reception          `gorm:"foreignKey:ReceptionID"`
	EmergencyReception []EmergencyReception `gorm:"foreignKey:EmergencyReceptionID"`
}

type PersonalInfo struct {
	gorm.Model

	PatientID      uint   `gorm:"not null;uniqueIndex" json:"patient_id" example:"1" rus:"ID пациента"`
	PassportSeries string `gorm:"not null" json:"passport_series" example:"4510 123456" rus:"Серия и номер паспорта"`
	SNILS          string `gorm:"unique;not null" json:"snils" example:"123-456-789 00" rus:"СНИЛС"`
	OMS            string `gorm:"unique;not null" json:"oms" example:"1234567890123456" rus:"Полис ОМС"`
}

type ContactInfo struct {
	gorm.Model

	PatientID uint   `gorm:"not null;uniqueIndex" json:"patient_id" example:"1" rus:"ID пациента"`
	Phone     string `gorm:"not null" json:"phone" example:"+79991234567" rus:"Телефон"`
	Email     string `gorm:"not null" json:"email" example:"patient@example.com" rus:"Email"`
	Address   string `gorm:"not null" json:"address" example:"Москва, ул. Пушкина, д. 10" rus:"Адрес"`
}

type Allergy struct {
	gorm.Model

	Name string `gorm:"not null" json:"name" example:"Пенициллин" rus:"Название"`
}

// TODO: Написать функцию получения аллергий для пациента
type PatientsAllergy struct {
	gorm.Model
	AllergyID uint    `gorm:"not null" json:"allergy_id" example:"1" rus:"ID"`
	Allergy   Allergy `gorm:"foreignKey:AllergyID" json:"-"`

	PatientID uint    `gorm:"not null" json:"patient_id" example:"1" rus:"ID"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"-"`

	Description string `json:"description" example:"Тяжелой степени"`
}

type ReceptionStatus string

const (
	StatusScheduled ReceptionStatus = "scheduled" // "Запланирован"
	StatusCompleted ReceptionStatus = "completed" // "Завершен"
	StatusCancelled ReceptionStatus = "cancelled" // "Отменен"
	StatusNoShow    ReceptionStatus = "no_show"   // "Не явился"
)

type EmergencyStatus string

const (
	EmergencyStatusScheduled ReceptionStatus = "scheduled" // "В ожидании"
	EmergencyStatusAccepted  ReceptionStatus = "accepted"  // "Принят"
	EmergencyStatusOnPlace   ReceptionStatus = "on_place"  // "На месте"
	EmergencyStatusCompleted ReceptionStatus = "completed" // "Завершен"
	EmergencyStatusCancelled ReceptionStatus = "cancelled" // "Отменен"
	EmergencyStatusNoShow    ReceptionStatus = "no_show"   // "Не явился"
)

type EmergencyReception struct {
	gorm.Model

	DoctorID        uint            `gorm:"index" json:"doctor_id" example:"1" rus:"ID врача"`
	PatientID       uint            `gorm:"not null;index" json:"patient_id" example:"1" rus:"ID пациента"`
	Status          EmergencyStatus `json:"status"`
	Priority        bool            `json:"priority"` // 1 - экстренный, 0 - неотложный
	Address         string          `gorm:"not null" json:"address" example:"Москва, ул. Ленина, д. 15" rus:"Адрес"`
	Date            time.Time       `gorm:"not null" json:"date" example:"2023-10-15T14:30:00Z" rus:"Дата приема"`
	Diagnosis       string          `json:"diagnosis" example:"ОРВИ" rus:"Диагноз"`
	Recommendations string          `json:"recommendations" example:"Постельный режим" rus:"Рекомендации"`
}

type Reception struct {
	gorm.Model

	DoctorID        uint            `gorm:"not null;index" json:"doctor_id" example:"1" rus:"ID врача"`
	PatientID       uint            `gorm:"not null;index" json:"patient_id" example:"1" rus:"ID пациента"`
	Date            time.Time       `gorm:"not null" json:"date" example:"2023-10-15T14:30:00Z" rus:"Дата приема"`
	Diagnosis       string          `json:"diagnosis" example:"ОРВИ" rus:"Диагноз"`
	Recommendations string          `json:"recommendations" example:"Постельный режим" rus:"Рекомендации"`
	IsOut           bool            `gorm:"not null;default:false" json:"is_out" example:"true" rus:"Вызов на выезд"` // 0 -  в стационаре, 1 - выезд
	Status          ReceptionStatus `gorm:"not null" json:"status" example:"scheduled" rus:"Статус"`
	Address         string          `gorm:"not null" json:"address" example:"Москва, ул. Ленина, д. 15" rus:"Адрес"`
}

type MedService struct {
	gorm.Model

	Name  string `gorm:"not null" json:"name" example:"EKG" rus:"ЭКГ"`
	Price uint   `gorm:"not null" json:"price" example:"100" rus:"Цена"`
}

// TODO: Написать функцию получения emergency_reception + получение med_service
type EmergencyReceptionMedServices struct {
	gorm.Model

	EmergencyReceptionID uint               `gorm:"not null" json:"emergency_reception_id" ` //
	EmergencyReception   EmergencyReception `gorm:"foreignKey:EmergencyReceptionID" json:"-"`

	MedServiceID uint       `gorm:"not null" json:"med_service_id" `
	MedService   MedService `gorm:"foreignKey:MedServiceID" json:"-"`
}
