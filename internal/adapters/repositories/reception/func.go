package receptionRepository

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (r *ReceptionRepositoryImpl) CreateReception(reception *entities.Reception) error {
	return r.db.Create(reception).Error
}

func (r *ReceptionRepositoryImpl) UpdateReception(reception *entities.Reception) error {
	return r.db.Save(reception).Error
}

func (r *ReceptionRepositoryImpl) DeleteReception(id uint) error {
	return r.db.Delete(&entities.Reception{}, id).Error
}

func (r *ReceptionRepositoryImpl) GetReceptionByID(id uint) (*entities.Reception, error) {
	var reception entities.Reception
	if err := r.db.First(&reception, id).Error; err != nil {
		return nil, err
	}
	return &reception, nil
}

func (r *ReceptionRepositoryImpl) GetReceptionByDoctorID(doctorID uint) ([]entities.Reception, error) {
	var receptions []entities.Reception
	err := r.db.Where("doctor_id = ?", doctorID).Find(&receptions).Error
	return receptions, err
}

func (r *ReceptionRepositoryImpl) GetReceptionByPatientID(patientID uint) ([]entities.Reception, error) {
	var receptions []entities.Reception
	err := r.db.Where("patient_id = ?", patientID).Find(&receptions).Error
	return receptions, err
}

func (r *ReceptionRepositoryImpl) GetReceptionByDateRange(start, end time.Time) ([]entities.Reception, error) {
	var receptions []entities.Reception
	err := r.db.Where("date BETWEEN ? AND ?", start, end).Find(&receptions).Error
	return receptions, err
}
