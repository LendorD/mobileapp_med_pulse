package emergencyReception

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
)

func (r *EmergencyReceptionRepositoryImpl) Create(er *entities.EmergencyReception) error {
	return r.db.Create(er).Error
}

func (r *EmergencyReceptionRepositoryImpl) Update(er *entities.EmergencyReception) error {
	return r.db.Save(er).Error
}

func (r *EmergencyReceptionRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.EmergencyReception{}, id).Error
}

func (r *EmergencyReceptionRepositoryImpl) GetByID(id uint) (*entities.EmergencyReception, error) {
	var er entities.EmergencyReception
	if err := r.db.First(&er, id).Error; err != nil {
		return nil, err
	}
	return &er, nil
}

func (r *EmergencyReceptionRepositoryImpl) GetByDoctorID(doctorID uint) ([]entities.EmergencyReception, error) {
	var list []entities.EmergencyReception
	err := r.db.Where("doctor_id = ?", doctorID).Find(&list).Error
	return list, err
}

func (r *EmergencyReceptionRepositoryImpl) GetByPatientID(patientID uint) ([]entities.EmergencyReception, error) {
	var list []entities.EmergencyReception
	err := r.db.Where("patient_id = ?", patientID).Find(&list).Error
	return list, err
}

func (r *EmergencyReceptionRepositoryImpl) GetByDateRange(start, end time.Time) ([]entities.EmergencyReception, error) {
	var list []entities.EmergencyReception
	err := r.db.Where("date BETWEEN ? AND ?", start, end).Find(&list).Error
	return list, err
}

func (r *EmergencyReceptionRepositoryImpl) GetPriorityCases() ([]entities.EmergencyReception, error) {
	var list []entities.EmergencyReception
	err := r.db.Where("priority = true").Find(&list).Error
	return list, err
}
