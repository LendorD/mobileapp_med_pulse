package patientsAllergy

import (
	_ "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type PatientsAllergyRepositoryImpl struct {
	db *gorm.DB
}

func NewPatientsAllergyRepository(db *gorm.DB) interfaces.PatientsAllergyRepository {
	return &PatientsAllergyRepositoryImpl{db: db}
}
