package entities

import "time"

type Emk struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	PatientID string     `gorm:"type:varchar(255);index" json:"patient_id"` // связь с пациентом
	CallID    string     `gorm:"uniqueIndex;type:varchar(255)" json:"call_id"`
	Status    CallStatus `gorm:"not null;default:'received'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	MedServices []byte `gorm:"type:jsonb" json:"med_services,omitempty"`

	Doctor   DoctorData `gorm:"foreignKey:ReceptionID" json:"doctor"`
	Flg      Flg        `gorm:"foreignKey:ReceptionID" json:"flg"`
	Analysis Analysis   `gorm:"foreignKey:ReceptionID" json:"analysis"`
}
