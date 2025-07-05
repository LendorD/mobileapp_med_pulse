package patientsAllergyRepository

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *PatientsAllergyRepositoryImpl) Create(pa *entities.PatientsAllergy) error {
	return r.db.Create(pa).Error
}

func (r *PatientsAllergyRepositoryImpl) Update(pa *entities.PatientsAllergy) error {
	return r.db.Save(pa).Error
}

func (r *PatientsAllergyRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.PatientsAllergy{}, id).Error
}

func (r *PatientsAllergyRepositoryImpl) GetByID(id uint) (*entities.PatientsAllergy, error) {
	var pa entities.PatientsAllergy
	if err := r.db.First(&pa, id).Error; err != nil {
		return nil, err
	}
	return &pa, nil
}

func (r *PatientsAllergyRepositoryImpl) GetByPatientID(patientID uint) ([]entities.PatientsAllergy, error) {
	var list []entities.PatientsAllergy
	if err := r.db.Where("patient_id = ?", patientID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
