package entities

import "gorm.io/gorm"

// TODO: Написать функцию получения emergency_reception + получение med_service
type EmergencyReceptionMedServices struct {
	gorm.Model
	Diagnosis       string `json:"diagnosis" example:"ОРВИ" rus:"Диагноз"`
	Recommendations string `json:"recommendations" example:"Постельный режим" rus:"Рекомендации"`

	EmergencyReceptionID uint               `gorm:"not null" json:"emergency_reception_id" ` //
	EmergencyReception   EmergencyReception `gorm:"foreignKey:EmergencyReceptionID" json:"-"`

	MedServiceID uint       `gorm:"not null" json:"med_service_id" `
	MedService   MedService `gorm:"foreignKey:MedServiceID" json:"-"`
}
