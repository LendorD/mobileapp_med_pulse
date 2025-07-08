package entities

import "gorm.io/gorm"

type PatientsAllergy struct {
	gorm.Model
	AllergyID uint    `gorm:"not null" json:"allergy_id" example:"1" rus:"ID"`
	Allergy   Allergy `gorm:"foreignKey:AllergyID" json:"-"`

	PatientID uint    `gorm:"not null" json:"patient_id" example:"1" rus:"ID"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"-"`

	Description string `json:"description" example:"Тяжелой степени"`
}
