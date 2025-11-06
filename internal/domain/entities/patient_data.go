package entities

// PatientData — данные пациента, связанные с вызовом
// type PatientData struct {
// 	ID uint `gorm:"primaryKey" json:"id"`

// 	// Идентификаторы
// 	PatientID string `gorm:"column:patient_id;type:varchar(255);uniqueIndex" json:"patient_id"`

// 	// Основная информация
// 	FullName    string `gorm:"type:varchar(255)" json:"full_name"`
// 	Age         int    `gorm:"type:int4" json:"age"`
// 	BirthDate   string `gorm:"type:varchar(10)" json:"birth_date"`
// 	Gender      bool
// 	MobilePhone string `gorm:"type:varchar(20)" json:"mobile_phone"`

// 	// Связи
// 	MedicalCard OneCMedicalCard `gorm:"foreignKey:PatientID;references:PatientID" json:"medical_card,omitempty"`
// 	Emk         []Emk           `gorm:"foreignKey:PatientID;references:PatientID" json:"emk,omitempty"`
// }
