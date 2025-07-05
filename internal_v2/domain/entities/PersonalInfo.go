package entities

import "gorm.io/gorm"

type PersonalInfo struct {
	gorm.Model

	PatientID      uint   `gorm:"not null;uniqueIndex" json:"patient_id" example:"1" rus:"ID пациента"`
	PassportSeries string `gorm:"not null" json:"passport_series" example:"4510 123456" rus:"Серия и номер паспорта"`
	SNILS          string `gorm:"unique;not null" json:"snils" example:"123-456-789 00" rus:"СНИЛС"`
	OMS            string `gorm:"unique;not null" json:"oms" example:"1234567890123456" rus:"Полис ОМС"`
}
