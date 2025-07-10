package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FullName       string `json:"full_name"`
	Specialization string `json:"specialization"`
	Login          string `json:"login"`
	PasswordHash   string `json:"-"`
}
