package models

import "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

// Call — зеркало entities.OneCReception для внешнего API (без полей БД)
type Call struct {
	CallID      string              `json:"call_id"`
	Address     string              `json:"address"`
	Phone       string              `json:"phone"`
	Status      entities.CallStatus `json:"status"`
	MedServices []byte              `json:"med_services,omitempty"`

	Patients   []entities.OneCPatientListItem `json:"patients,omitempty"`
	Doctor     entities.DoctorData            `json:"doctor"`
	Receptions []entities.Receptions          `json:"receptions,omitempty"`
	Flg        entities.Flg                   `json:"flg"`
	Analysis   entities.Analysis              `json:"analysis"`
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
