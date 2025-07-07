package reception

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type ReceptionRepositoryImpl struct {
	db *gorm.DB
}

func NewReceptionRepository(db *gorm.DB) interfaces.ReceptionRepository {
	return &ReceptionRepositoryImpl{db: db}
}
