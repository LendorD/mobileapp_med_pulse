package models

import "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

// EmergencyReceptionMedServicesResponse - информация о медицинских услугах при срочном приеме
// @Description Содержит данные пациента, диагноз и рекомендации по срочному приему
type EmergencyReceptionMedServicesResponse struct {
	Patient         entities.Patient `json:"patient"`         // Данные пациента
	Diagnosis       string           `json:"diagnosis"`       // Поставленный диагноз
	Recommendations string           `json:"recommendations"` // Рекомендации врача
}
