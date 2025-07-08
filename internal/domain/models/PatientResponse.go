package models

import (
	"time"

	"gorm.io/gorm"
)

type UpdatePatientRequest struct {
	ID        uint
	FullName  string
	BirthDate time.Time
}

type CreatePatientRequest struct {
	FullName  string
	BirthDate time.Time
	IsMale    bool
}

type ShortPatientResponse struct {
	gorm.Model
	FullName  string    `json:"-" example:"Смирнов Алексей Петрович" rus:"ФИО"`
	BirthDate time.Time `json:"birth_date" example:"1980-05-15T00:00:00Z" rus:"Дата рождения"`
	IsMale    bool      `json:"is_male" example:"true" rus:"Пол (true - мужской)"`
}
