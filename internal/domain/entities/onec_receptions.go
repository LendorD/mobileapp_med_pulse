package entities

import "time"

// OneCReception — приём (список пациентов) по вызову скорой
type OneCReception struct {
	ID        uint   `gorm:"primaryKey"`
	CallID    string `gorm:"uniqueIndex"` // ID вызова из 1С
	Status    string `gorm:"not null"`    // "received", "edited", "pending_sync", "synced", "sync_failed"
	Data      []byte `gorm:"type:jsonb"`  // Вся структура от 1С (включая пациента, услуги и т.д.)
	CreatedAt time.Time
	UpdatedAt time.Time
}
