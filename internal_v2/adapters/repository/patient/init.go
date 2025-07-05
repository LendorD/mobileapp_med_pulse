package patient

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type PatientRepositoryImpl struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) interfaces.PatientRepository {
	return &PatientRepositoryImpl{db: db}
}
