package contactInfo

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type ContactInfoRepositoryImpl struct {
	db *gorm.DB
}

func NewContactInfoRepository(db *gorm.DB) interfaces.ContactInfoRepository {
	return &ContactInfoRepositoryImpl{db: db}
}
