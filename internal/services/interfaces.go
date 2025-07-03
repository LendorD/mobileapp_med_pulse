package services

import "github.com/AlexanderMorozov1919/mobileapp/internal/models"

type SmpService interface {
	GetCallings(doctorID uint) ([]ReceptionResponce, error)
}

type ReceptionService interface {
}

type PatientService interface {
	GetAllPatientsByDoctorID(doctorID uint) ([]models.ShortPatientResponse, error)
}
