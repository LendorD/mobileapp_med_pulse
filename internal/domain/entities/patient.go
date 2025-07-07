package entities

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FullName  string    `gorm:"not null" json:"-" example:"Смирнов Алексей Петрович" rus:"ФИО"`
	BirthDate time.Time `gorm:"not null" json:"birth_date" example:"1980-05-15T00:00:00Z" rus:"Дата рождения"`
	IsMale    bool      `gorm:"not null" json:"is_male" example:"true" rus:"Пол (true - мужской)"`

	Reception          []Reception          `gorm:"foreignKey:ReceptionID"`
	EmergencyReception []EmergencyReception `gorm:"foreignKey:EmergencyReceptionID"`
}
