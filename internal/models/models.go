package models

import "time"

type Doctor struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName      string    `gorm:"size:100;not null" json:"first_name"`
	MiddleName     string    `gorm:"size:100" json:"middle_name"`
	LastName       string    `gorm:"size:100;not null" json:"last_name"`
	Login          string    `gorm:"size:50;unique;not null" json:"login"`
	PasswordHash   string    `gorm:"size:255;not null" json:"password_hash"`
	Specialization string    `gorm:"size:150;not null" json:"specialization"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Patient struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName  string    `gorm:"size:100;not null" json:"first_name"`
	MiddleName string    `gorm:"size:100;not null" json:"middle_name"`
	LastName   string    `gorm:"size:100;not null" json:"last_name"`
	FullName   string    `gorm:"size:300;not null" json:"full_name"`
	BirthDate  time.Time `gorm:"not null" json:"birth_date"`  // Лучше использовать time.Time вместо int
	IsMale     bool      `gorm:"not null" json:"is_male"`     // true - мужской, false - женский
	SNILS      string    `gorm:"size:14;unique" json:"snils"` // СНИЛС
	OMS        string    `gorm:"size:16;unique" json:"oms"`   // Полис ОМС
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ContactInfo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID uint      `gorm:"not null;index" json:"patient_id"`
	Phone     string    `gorm:"size:20;not null" json:"phone"` // Строка лучше для телефонных номеров
	Email     string    `gorm:"size:100" json:"email"`
	Address   string    `gorm:"type:text;not null" json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ReceptionStatus string

const (
	StatusScheduled ReceptionStatus = "scheduled" // Запланирован
	StatusCompleted ReceptionStatus = "completed" // Завершен
	StatusCancelled ReceptionStatus = "cancelled" // Отменен
	StatusNoShow    ReceptionStatus = "no_show"   // Не явился
)

type Reception struct {
	ID              uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	DoctorID        uint            `gorm:"not null;index" json:"doctor_id"`
	PatientID       uint            `gorm:"not null;index" json:"patient_id"`
	Date            time.Time       `gorm:"not null" json:"date"`
	Diagnosis       string          `gorm:"type:text" json:"diagnosis"`       // Диагноз
	Recommendations string          `gorm:"type:text" json:"recommendations"` // Рекомендации
	IsSMP           bool            `gorm:"default:false" json:"is_smp"`      // Работает в СМП (скорая медицинская помощь): true - да, false - нет
	Status          ReceptionStatus `gorm:"type:varchar(20);not null;default:'scheduled'" json:"status"`
	Address         string          `gorm:"type:text" json:"address"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

type Allergy struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID   uint      `gorm:"not null;index" json:"patient_id"`
	Name        string    `gorm:"size:150;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	RecordedAt  time.Time `gorm:"autoCreateTime" json:"recorded_at"`
}

type PatientCard struct {
	Patient     Patient     `json:"patient"`
	ContactInfo ContactInfo `json:"contact_info"`
	Receptions  []Reception `json:"receptions"` // Массив (слайс) структур Receptions
	Allergies   []Allergy   `json:"allergies"`  // Массив структур аллергий
}
