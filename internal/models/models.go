package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	FirstName      string `gorm:"size:100;not null" json:"first_name"`
	MiddleName     string `gorm:"size:100" json:"middle_name"`
	Surname        string `gorm:"size:100;not null" json:"surname"`
	Login          string `gorm:"size:50;unique;not null" json:"login"`
	PasswordHash   string `gorm:"size:255;not null" json:"-"`
	Specialization string `gorm:"size:100;not null" json:"specialization"`
}

type Patient struct {
	gorm.Model
	FirstName  string    `gorm:"size:100;not null" json:"first_name"`
	Surname    string    `gorm:"size:100;not null" json:"surname"`
	MiddleName string    `gorm:"size:100" json:"middle_name"`
	FullName   string    `gorm:"size:300" json:"full_name"`
	BirthDate  time.Time `gorm:"not null" json:"birth_date"`
	IsMale     bool      `gorm:"not null" json:"is_male"`
	SNILS      string    `gorm:"size:14;unique" json:"snils"`
	OMS        string    `gorm:"size:16;unique" json:"oms"`
}

type ContactInfo struct {
	gorm.Model
	PatientID uint   `gorm:"not null" json:"patient_id"`
	Phone     string `gorm:"size:20;not null" json:"phone"`
	Email     string `gorm:"size:100" json:"email"`
	Address   string `gorm:"size:200" json:"address"`
}

type ReceptionStatus string

const (
	StatusScheduled ReceptionStatus = "scheduled"
	StatusCompleted ReceptionStatus = "completed"
	StatusCancelled ReceptionStatus = "cancelled"
	StatusNoShow    ReceptionStatus = "no_show"
)

type Reception struct {
	gorm.Model
	DoctorID        uint            `gorm:"not null" json:"doctor_id"`
	PatientID       uint            `gorm:"not null" json:"patient_id"`
	Date            time.Time       `gorm:"not null" json:"date"`
	Diagnosis       string          `gorm:"type:text" json:"diagnosis"`
	Recommendations string          `gorm:"type:text" json:"recommendations"`
	IsSMP           bool            `gorm:"default:false" json:"is_smp"`
	Status          ReceptionStatus `gorm:"type:varchar(20);default:'scheduled'" json:"status"`
}

type Allergy struct {
	gorm.Model
	PatientID   uint   `gorm:"not null" json:"patient_id"`
	Name        string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

type PatientCard struct {
	Patient     Patient     `json:"patient"`
	ContactInfo ContactInfo `json:"contact_info"`
	Receptions  []Reception `json:"receptions"`
	Allergies   []Allergy   `json:"allergies"`
}
