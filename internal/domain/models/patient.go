package models

import (
	"gorm.io/gorm"
	"time"
)

type UpdatePatientRequest struct {
	ID        uint
	FullName  string
	BirthDate time.Time
}

type CreatePatientRequest struct {
	FullName  string    `json:"full_name" example:"Смирнов Алексей Петрович" rus:"ФИО"`
	BirthDate time.Time `json:"birth_date" example:"1980-05-15T00:00:00Z" rus:"Дата рождения"`
	IsMale    bool      `json:"is_male" example:"true" rus:"Пол (true - мужской)"`
}

type ShortPatientResponse struct {
	gorm.Model
	FullName  string    `json:"full_name" example:"Смирнов Алексей Петрович" rus:"ФИО"`
	BirthDate time.Time `json:"birth_date" example:"1980-05-15T00:00:00Z" rus:"Дата рождения"`
	IsMale    bool      `json:"is_male" example:"true" rus:"Пол (true - мужской)"`
}

type PatientResponse struct {
}
