package personalinfo

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"gorm.io/gorm"
)

type PersonalInfoRepositoryImpl struct {
	db *gorm.DB
}

func NewPersonalInfoRepository(db *gorm.DB) interfaces.PersonalInfoRepository {
	return &PersonalInfoRepositoryImpl{db: db}
}
