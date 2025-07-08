package doctor

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type DoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) interfaces.DoctorRepository {
	return &DoctorRepository{db: db}
}
