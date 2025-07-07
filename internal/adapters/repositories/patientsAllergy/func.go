package patientsAllergy

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *PatientsAllergyRepositoryImpl) CreatePatientsAllergy(pa *entities.PatientsAllergy) error {
	return r.db.Create(pa).Error
}

func (r *PatientsAllergyRepositoryImpl) UpdatePatientsAllergy(pa *entities.PatientsAllergy) error {
	return r.db.Save(pa).Error
}

func (r *PatientsAllergyRepositoryImpl) DeletePatientsAllergy(id uint) error {
	return r.db.Delete(&entities.PatientsAllergy{}, id).Error
}

func (r *PatientsAllergyRepositoryImpl) ExistsAllergy(patientID, allergyID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entities.PatientsAllergy{}).
		Where("patient_id = ? AND allergy_id = ?", patientID, allergyID).
		Count(&count).Error

	return count > 0, err
}

func (r *PatientsAllergyRepositoryImpl) GetPatientAllergyByPatientID(patientID uint) (*entities.PatientsAllergy, error) {
	var patient *entities.PatientsAllergy
	if err := r.db.Where("patient_id = ?", patientID).Find(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func (r *PatientsAllergyRepositoryImpl) GetPatientsAllergiesByPatientID(patientID uint) ([]entities.PatientsAllergy, error) {
	//TODO implement me
	panic("implement me")
}
func (r *PatientsAllergyRepositoryImpl) GetAllergyByPatientID(patientID uint) ([]entities.PatientsAllergy, error) {
	//TODO implement me
	panic("implement me")
}
