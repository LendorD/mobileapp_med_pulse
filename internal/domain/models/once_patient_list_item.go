package models

type PatientListUpdate struct {
	Patients []PatientListItem `json:"patients"`
}

// PatientListItem — краткая информация о пациенте для списка
type PatientListItem struct {
	PatientID string `json:"patient_id"`
	FullName  string `json:"full_name"`  // ФИО
	Gender    bool   `json:"gender"`     // Пол: true — мужской, false — женский
	BirthDate string `json:"birth_date"` // Дата рождения
}
