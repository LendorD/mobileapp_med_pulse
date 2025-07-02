package repository

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
)

type DoctorRepository interface {
	Create(doctor *models.Doctor) error
	Update(docor *models.Doctor) error
	Delete(id uint) error
	GetByID(id uint) (*models.Doctor, error)
	GetName(id uint) (string, error)
	GetSpecialization(id uint) (string, error)
	GetPassHash(id uint) (string, error)
}

type ReceptionRepository interface {
	Create(reception *models.Reception) error
	Update(reception *models.Reception) error
	Delete(id uint) error
	GetByID(id uint) (*models.Reception, error)
	GetAllByDoctorID(doctorID uint) ([]models.Reception, error)
	GetAllByDate(date time.Time) ([]models.Reception, error)
}
