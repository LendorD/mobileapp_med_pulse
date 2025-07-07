package contactInfo

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *ContactInfoRepositoryImpl) Create(info *entities.ContactInfo) error {
	return r.db.Create(info).Error
}

func (r *ContactInfoRepositoryImpl) Update(info *entities.ContactInfo) error {
	return r.db.Save(info).Error
}

func (r *ContactInfoRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.ContactInfo{}, id).Error
}

func (r *ContactInfoRepositoryImpl) GetByID(id uint) (*entities.ContactInfo, error) {
	var info entities.ContactInfo
	if err := r.db.First(&info, id).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *ContactInfoRepositoryImpl) GetByPatientID(patientID uint) (*entities.ContactInfo, error) {
	var info entities.ContactInfo
	if err := r.db.Where("patient_id = ?", patientID).First(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}
