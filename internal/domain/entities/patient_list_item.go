package entities

type PatientListUpdate struct {
	Patients []OneCPatientListItem `json:"patients"`
}

// PatientListItem — краткая информация о пациенте для списка
type OneCPatientListItem struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	PatientID string `gorm:"not null;uniqueIndex" json:"patient_id"`
	FullName  string `gorm:"not null" json:"full_name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}
