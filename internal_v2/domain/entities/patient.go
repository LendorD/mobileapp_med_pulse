package entities

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	IsMale    bool      `json:"is_male"`
}
