package handlers

import "github.com/AlexanderMorozov1919/mobileapp/internal/services"

type PatientHandler struct {
	patientService services.PatientService
}

func NewPatientHandler(patientService services.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}
