package doctor

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

func (r *DoctorRepository) GetDoctorByID(id uint) (entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.
		Preload("Specialization").
		First(&doctor, id).
		Error; err != nil {
		return entities.Doctor{}, errors.NewDBError("Error Get Doctor By Id", err)
	}
	return doctor, nil
}

func (r *DoctorRepository) GetDoctorByLogin(login string) (entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.Where("login = ?", login).First(&doctor).Error; err != nil {
		return entities.Doctor{}, errors.NewDBError("Error Get Doctor By Login", err)
	}
	return doctor, nil
}
