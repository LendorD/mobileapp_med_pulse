package entities

import "time"

type File struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	MinioKey  string `gorm:"not null;uniqueIndex"` // "minio_id"
	Filename  string `gorm:"not null"`
	CreatedAt time.Time
}
