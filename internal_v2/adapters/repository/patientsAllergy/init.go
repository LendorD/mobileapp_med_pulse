package patientsAllergyRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type PatientsAllergyRepositoryImpl struct {
	db *gorm.DB
}

func NewPatientsAllergyRepository(db *gorm.DB) interfaces.PatientsAllergyRepository {
	return &PatientsAllergyRepositoryImpl{db: db}
}
