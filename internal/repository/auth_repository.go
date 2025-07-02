package repository

import (
	"gorm.io/gorm"

	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateDoctor(doctor *models.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *AuthRepository) FindDoctorByLogin(login string) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Where("login = ?", login).First(&doctor).Error
	return &doctor, err
}
