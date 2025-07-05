package personalinfo

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *PersonalInfoRepositoryImpl) Create(info *entities.PersonalInfo) error {
	return r.db.Create(info).Error
}

func (r *PersonalInfoRepositoryImpl) Update(info *entities.PersonalInfo) error {
	return r.db.Save(info).Error
}

func (r *PersonalInfoRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.PersonalInfo{}, id).Error
}

func (r *PersonalInfoRepositoryImpl) GetByID(id uint) (*entities.PersonalInfo, error) {
	var info entities.PersonalInfo
	if err := r.db.First(&info, id).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *PersonalInfoRepositoryImpl) GetByPatientID(patientID uint) (*entities.PersonalInfo, error) {
	var info entities.PersonalInfo
	if err := r.db.Where("patient_id = ?", patientID).First(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}
