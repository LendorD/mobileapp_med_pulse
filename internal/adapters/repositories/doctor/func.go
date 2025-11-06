package doctor

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

func (r *DoctorRepository) GetDoctorByID(ctx context.Context, id uint) (entities.DoctorData, error) {
	var doctor entities.DoctorData
	db := r.db.GetDB(ctx)
	if err := db.
		Preload("Specialization").
		First(&doctor, id).
		Error; err != nil {
		return entities.DoctorData{}, errors.NewDBError("Error Get Doctor By Id", err)
	}
	return doctor, nil
}

func (r *DoctorRepository) GetDoctorByLogin(ctx context.Context, login string) (entities.DoctorData, error) {
	var doctor entities.DoctorData
	db := r.db.GetDB(ctx)
	if err := db.Where("login = ?", login).First(&doctor).Error; err != nil {
		return entities.DoctorData{}, errors.NewDBError("Error Get Doctor By Login", err)
	}
	return doctor, nil
}
