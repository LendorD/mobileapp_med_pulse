package models

import (
	"time"
)

type UpdatePatientRequest struct {
	ID        uint
	FullName  string
	BirthDate time.Time
}

type CreatePatientRequest struct {
	FullName  string
	BirthDate time.Time
	IsMale    bool
}
