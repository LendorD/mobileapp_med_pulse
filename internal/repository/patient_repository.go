package repository

import (
	"errors"
	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	GetByID(id uint) (*models.Patient, error)
	SearchByName(name string) ([]models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
	List(limit, offset int) ([]models.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) GetByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) GetBySNILS(snils string) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Where("snils = ?", snils).First(&patient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) GetByOMS(oms string) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Where("oms = ?", oms).First(&patient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) SearchByName(name string) ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Where("full_name LIKE ? OR first_name LIKE ? OR surname LIKE ?",
		"%"+name+"%", "%"+name+"%", "%"+name+"%").
		Limit(50).
		Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *patientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}

func (r *patientRepository) List(limit, offset int) ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Limit(limit).Offset(offset).Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}
