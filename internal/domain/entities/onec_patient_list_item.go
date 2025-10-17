package entities

type PatientListUpdate struct {
	Patients []OneCPatientListItem `json:"patients"`
}

// PatientListItem — краткая информация о пациенте для списка
type OneCPatientListItem struct {
	ID        uint   `gorm:"primaryKey"`
	PatientID string `gorm:"not null;index"`
	FullName  string `gorm:"not null"`
	Gender    bool   // true — мужской
	BirthDate string // в формате "YYYY-MM-DD"
}
