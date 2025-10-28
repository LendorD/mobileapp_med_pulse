package models

import "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

// Call — основная структура вызова из 1С
type Call struct {
	CallID      string              `json:"call_id"`
	Address     string              `json:"address"` // Адрес (адрес вызова)
	Phone       string              `json:"phone"`
	Status      entities.CallStatus // "received", "edited", "pending_sync", "synced", "sync_failed"
	Patient     []entities.PatientData
	Doctor      entities.DoctorData
	Receptions  []entities.Receptions // Заключения врача
	Flg         entities.Flg
	Analysis    entities.Analysis
	MedServices []byte `gorm:"type:jsonb"`
}

// Patient — данные пациента
type Patient struct {
	FullName    string               `json:"full_name"`   // ФИО
	BirthDate   string               `json:"birth_date"`  // Дата рождения
	Age         string               `json:"age"`         // Возраст
	Gender      bool                 `json:"gender"`      // Пол: true — мужской, false — женский
	Phone       string               `json:"phone"`       // Телефон
	Snils       string               `json:"snils"`       // СНИЛС
	Policy      entities.Policy      `json:"policy"`      // Полис
	Certificate entities.Certificate `json:"certificate"` // Сертификат
}
