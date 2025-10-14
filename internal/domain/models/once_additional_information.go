package models

// Policy — данные полиса
type Policy struct {
	Type       string `json:"type"`        // Вид полиса
	Number     string `json:"number"`      // Номер полиса
	ExpiryDate string `json:"expiry_date"` // Срок действия
	Contractor string `json:"contractor"`  // Контрагент
}

// Certificate — данные сертификата
type Certificate struct {
	Type       string `json:"type"`        // Вид сертификата
	StartDate  string `json:"start_date"`  // Дата начала действия
	ExpiryDays string `json:"expiry_days"` // Дата окончания дней
}
