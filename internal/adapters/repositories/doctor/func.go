package doctor

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

func (r *DoctorRepository) GetDoctorByID(ctx context.Context, id uint) (entities.Doctor, error) {
	var doctor entities.Doctor
	db := r.db.GetDB(ctx)
	if err := db.
		Preload("Specialization").
		First(&doctor, id).
		Error; err != nil {
		return entities.Doctor{}, errors.NewDBError("Error Get Doctor By Id", err)
	}
	return doctor, nil
}

func (r *DoctorRepository) GetDoctorByLogin(ctx context.Context, login string) (entities.Doctor, error) {
	var doctor entities.Doctor
	db := r.db.GetDB(ctx)
	if err := db.Where("login = ?", login).First(&doctor).Error; err != nil {
		return entities.Doctor{}, errors.NewDBError("Error Get Doctor By Login", err)
	}
	return doctor, nil
}
