package patient

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type PatientRepositoryImpl struct {
	db *base.BaseRepository
}

func NewPatientRepository(db *gorm.DB) interfaces.PatientRepository {
	return &PatientRepositoryImpl{db: base.NewBaseRepository(db)}
}
