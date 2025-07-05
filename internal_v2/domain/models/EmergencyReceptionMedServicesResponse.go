package models

import "github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"

type EmergencyReceptionMedServicesResponse struct {
	Patient         entities.Patient `json:"patient"`
	Diagnosis       string           `json:"diagnosis"`
	Recommendations string           `json:"recommendations"`
}
