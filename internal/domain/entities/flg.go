package entities

import "time"

// Flg — данные ФЛГ, привязанные к вызову (1:1)
type Flg struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	ReceptionID uint `gorm:"uniqueIndex;not null" json:"-"`

	CreatedAt    time.Time `json:"created_at"`
	PatientID    uint      `json:"patient_id"`
	Organization string    `json:"organization"`
	Number       string    `json:"number"`
	Result       string    `json:"result"`
	Date         time.Time `json:"date"`
	FileID       *uint     `json:"file_id"`
	File         *File     `gorm:"foreignKey:FileID" json:"file,omitempty"`
}
