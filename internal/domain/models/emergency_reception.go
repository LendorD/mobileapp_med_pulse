package models

// EmergencyCallShortResponse - краткая информация о срочном приеме
// @Description Сокращенная информация о срочном приеме пациента
type EmergencyCallShortResponse struct {
	Id          uint   `json:"id" example:"3"`
	CreatedAt   string `json:"created_at" example:"2023-05-15T14:30:00Z"` // Дата и время создания
	Emergency   bool   `json:"emergency" example:"true"`
	Priority    *uint  `json:"priority" example:"1"`
	Address     string `json:"address" example:"ул. Ленина, д. 5, кв. 12"` // Адрес вызова
	Phone       string `json:"phone" example:"+79991234567"`               // Телефон для связи
	Description string `json:"description" example:"Пациент жалуется на сильную боль в груди"`
}

// CreateEmergencyCallRequest - Информация для создания вызова на скорой
// @Description Информация для создания вызова на скорой
type CreateEmergencyCallRequest struct {
	DoctorID    uint   `json:"doctor_id" example:"1"`
	Emergency   bool   `json:"emergency" example:"true"`
	Address     string `json:"address" example:"ул. Ленина, д. 5, кв. 12"` // Адрес вызова
	Phone       string `json:"phone" example:"+79991234567"`               // Телефон для связи
	Description string `json:"description" example:"Пациент жалуется на сильную боль в груди"`
}
