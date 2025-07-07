package allergyRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type AllergyRepositoryImpl struct {
	db *gorm.DB
}

func NewAllergyRepository(db *gorm.DB) interfaces.AllergyRepository {
	return &AllergyRepositoryImpl{db: db}
}
