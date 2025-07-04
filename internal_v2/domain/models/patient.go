package models

import "time"

type PatientRequest struct {
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	IsMale    bool      `json:"is_male"`
}
