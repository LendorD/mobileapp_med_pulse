package models

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Reception struct {
	DoctorID        uint                     `json:"doctor_id" example:"1" rus:"ID врача"`
	PatientID       uint                     `json:"patient_id" example:"1" rus:"ID пациента"`
	Date            time.Time                `json:"date" example:"2023-10-15T14:30:00Z" rus:"Дата приема"`
	Diagnosis       string                   `example:"ОРВИ" rus:"Диагноз"`
	Recommendations string                   `json:"recommendations" example:"Постельный режим" rus:"Рекомендации"`
	IsOut           bool                     `json:"is_out" example:"true" rus:"Вызов на выезд"` // 0 -  в стационаре, 1 - выезд
	Status          entities.ReceptionStatus `json:"status" example:"scheduled" rus:"Статус"`
	Address         string                   `json:"address" example:"Москва, ул. Ленина, д. 15" rus:"Адрес"`
}

type ReceptionResponse struct {
	DoctorID        uint                     `json:"doctor_id" example:"1" rus:"ID врача"`
	PatientID       uint                     `json:"patient_id" example:"1" rus:"ID пациента"`
	Date            time.Time                `json:"date" example:"2023-10-15T14:30:00Z" rus:"Дата приема"`
	Diagnosis       string                   `example:"ОРВИ" rus:"Диагноз"`
	Recommendations string                   `json:"recommendations" example:"Постельный режим" rus:"Рекомендации"`
	IsOut           bool                     `json:"is_out" example:"true" rus:"Вызов на выезд"` // 0 -  в стационаре, 1 - выезд
	Status          entities.ReceptionStatus `json:"status" example:"scheduled" rus:"Статус"`
	Address         string                   `json:"address" example:"Москва, ул. Ленина, д. 15" rus:"Адрес"`
}
