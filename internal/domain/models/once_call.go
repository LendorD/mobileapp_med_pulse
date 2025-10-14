package models

// Call — основная структура вызова из 1С
type Call struct {
	CallID       string     `json:"call_id"`
	Address      string     `json:"address"`       // Адрес (адрес вызова)
	Phone        string     `json:"phone"`         // Телефон пациента
	PatientCount int        `json:"patient_count"` // Кол-во пациентов
	Status       CallStatus `json:"status"`        // Статус вызова
	Patients     []Patient  `json:"patients"`      // Данные пациентов
}

// CallStatus — статус вызова
type CallStatus string

const (
	CallStatusCompleted CallStatus = "Выполнен"
	CallStatusWork      CallStatus = "В работе"
)

// Patient — данные пациента
type Patient struct {
	FullName    string      `json:"full_name"`   // ФИО
	BirthDate   string      `json:"birth_date"`  // Дата рождения
	Age         string      `json:"age"`         // Возраст
	Gender      bool        `json:"gender"`      // Пол: true — мужской, false — женский
	Phone       string      `json:"phone"`       // Телефон
	Snils       string      `json:"snils"`       // СНИЛС
	Policy      Policy      `json:"policy"`      // Полис
	Certificate Certificate `json:"certificate"` // Сертификат
}
