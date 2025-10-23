package entities

import "time"

type Flg struct {
	ID        uint      `gorm:"primarykey" json:"id" example:"1"`
	CreatedAt time.Time `json:"-"`

	PatientID    uint      `gorm:"index" json:"patient_id"`
	Organization string    `gorm:"not null" json:"organization" example:"Stavropol"`
	Number       string    `gorm:"not null" json:"number" example:"984212"`
	Result       string    `gorm:"not null" json:"result" example:"COVID"`
	Date         time.Time `json:"date" example:"2023-10-15T14:30:00Z"`

	FileID uint  `gorm:"index" json:"file_id"`
	File   *File `gorm:"foreignKey:FileID" json:"file,omitempty"`
}
