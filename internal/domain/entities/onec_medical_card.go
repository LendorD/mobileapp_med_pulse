// internal/domain/entities/onec_medical_card.go
package entities

// OneCMedicalCard — мед карта пациента,  используется и для БД, и для JSON
type OneCMedicalCard struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	PatientID string `gorm:"not null;uniqueIndex" json:"patient_id"`

	// Основные поля
	DisplayName     string `gorm:"type:text" json:"display_name"`
	Age             string `gorm:"type:varchar(50)" json:"age"`
	BirthDate       string `gorm:"type:varchar(20)" json:"birth_date"`
	MobilePhone     string `gorm:"type:varchar(20)" json:"mobile_phone"`
	AdditionalPhone string `gorm:"type:varchar(20)" json:"additional_phone"`
	Address         string `gorm:"type:text" json:"address"`
	Email           string `gorm:"type:varchar(100)" json:"email"`
	Workplace       string `gorm:"type:text" json:"workplace"`
	Snils           string `gorm:"type:varchar(20);index" json:"snils"`

	// Вложенные структуры
	LegalRepresentative ClientRef   `gorm:"embedded;embeddedPrefix:legal_rep_" json:"legal_representative,omitempty"`
	Relative            *Relative   `gorm:"embedded;embeddedPrefix:relative_" json:"relative,omitempty"`
	AttendingDoctor     Doctor      `gorm:"embedded;embeddedPrefix:doctor_" json:"attending_doctor"`
	Policy              Policy      `gorm:"embedded;embeddedPrefix:policy_" json:"policy"`
	Certificate         Certificate `gorm:"embedded;embeddedPrefix:cert_" json:"certificate"`
}

type ClientRef struct {
	ID   string `gorm:"type:varchar(50)" json:"id"`
	Name string `gorm:"type:text" json:"name"`
}

type Relative struct {
	Status string `gorm:"type:varchar(100)" json:"status"`
	Name   string `gorm:"type:text" json:"name"`
}

type Doctor struct {
	FullName           string `gorm:"type:text" json:"full_name"`
	PolicyOrCertNumber string `gorm:"type:varchar(100)" json:"policy_or_cert_number"`
	AttachmentStart    string `gorm:"type:varchar(20)" json:"attachment_start"`
	AttachmentEnd      string `gorm:"type:varchar(20)" json:"attachment_end"`
	Clinic             string `gorm:"type:text" json:"clinic"`
}

type Policy struct {
	Number string `gorm:"type:varchar(50);index" json:"number"`
	Type   string `gorm:"type:varchar(50)" json:"type"`
}

type Certificate struct {
	Number string `gorm:"type:varchar(50)" json:"number"`
	Date   string `gorm:"type:varchar(20)" json:"date"`
}
