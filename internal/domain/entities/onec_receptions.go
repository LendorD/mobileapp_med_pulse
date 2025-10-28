package entities

import (
	"time"
)

// OneCReception — заявка на скорую
type OneCReception struct {
	ID          uint   `gorm:"primaryKey"`
	CallID      string `gorm:"uniqueIndex"` // ID вызова из 1С
	Address     string
	Phone       string
	Status      CallStatus `gorm:"not null"` // "received", "edited", "pending_sync", "synced", "sync_failed"
	Patient     []PatientData
	Doctor      DoctorData
	Receptions  []Receptions // Заключения врача
	Flg         Flg
	Analysis    Analysis
	MedServices []byte `gorm:"type:jsonb"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CallStatus — статус вызова
type CallStatus string

const (
	CallStatusCompleted CallStatus = "compleated"
	CallStatusWork      CallStatus = "proccess"
)

type Receptions struct {
	Data []byte `gorm:"type:jsonb"`
}
