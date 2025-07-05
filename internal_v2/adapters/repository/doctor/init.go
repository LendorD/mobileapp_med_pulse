package Allergy

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type DoctorRepositor struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) interfaces.DoctorRepository {
	return &DoctorRepositor{db: db}
}
