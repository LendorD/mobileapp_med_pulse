package entities

import "gorm.io/gorm"

type ContactInfo struct {
	gorm.Model

	PatientID uint   `gorm:"not null;uniqueIndex" json:"patient_id" example:"1" rus:"ID пациента"`
	Phone     string `gorm:"not null" json:"phone" example:"+79991234567" rus:"Телефон"`
	Email     string `gorm:"not null" json:"email" example:"patient@example.com" rus:"Email"`
	Address   string `gorm:"not null" json:"address" example:"Москва, ул. Пушкина, д. 10" rus:"Адрес"`
}
