package entities

import (
	"time"
)

// PersonalInfo представляет персональную информацию
type PersonalInfo struct {
	ID        uint      `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	PatientID      uint   `gorm:"not null;uniqueIndex" json:"patient_id" example:"1" rus:"ID пациента"`
	PassportSeries string `json:"passport_series" example:"4510 123456" rus:"Серия и номер паспорта"`
	SNILS          string `json:"snils" example:"123-456-789 00" rus:"СНИЛС"`
	OMS            string `json:"oms" example:"1234567890123456" rus:"Полис ОМС"`
}
