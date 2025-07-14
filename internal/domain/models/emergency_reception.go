package models

// EmergencyReceptionShortResponse - краткая информация о срочном приеме
// @Description Сокращенная информация о срочном приеме пациента
type EmergencyReceptionShortResponse struct {
	Date        string `json:"date" example:"2023-05-15T14:30:00Z"`        // Дата и время приема
	Status      string `json:"status" example:"завершен"`                  // Статус приема
	PatientName string `json:"patient_name" example:"Иванов Иван"`         // ФИО пациента
	Priority    bool   `json:"priority" example:"true"`                    // Приоритетный прием
	Address     string `json:"address" example:"ул. Ленина, д. 5, кв. 12"` // Адрес вызова
}
