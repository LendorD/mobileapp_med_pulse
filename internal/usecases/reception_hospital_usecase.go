package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type ReceptionHospitalUsecase struct {
	repo interfaces.ReceptionHospitalRepository
}

func NewReceptionHospitalUsecase(repo interfaces.ReceptionHospitalRepository) interfaces.ReceptionHospitalUsecase {
	return &ReceptionHospitalUsecase{
		repo: repo,
	}
}

func (u *ReceptionHospitalUsecase) GetReceptionsHospitalByPatientID(patientId uint) ([]entities.ReceptionHospital, *errors.AppError) {
	if patientId == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get patient",
			errors.ErrEmptyData,
			true,
		)
	}

	receptions, err := u.repo.GetReceptionHospitalByPatientID(patientId)
	if err != nil {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get receptions",
			errors.ErrEmptyData,
			true,
		)
	}

	if receptions == nil {
		return []entities.ReceptionHospital{}, nil
	}

	return receptions, nil
}

func (u *ReceptionHospitalUsecase) GetPatientsByDoctorID(doctorID uint, limit, offset int) ([]entities.Patient, *errors.AppError) {
	if doctorID == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get patient",
			errors.ErrEmptyData,
			true,
		)
	}

	patients, err := u.repo.GetPatientsByDoctorID(doctorID, limit, offset)
	if err != nil {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get receptions",
			errors.ErrEmptyData,
			true,
		)
	}

	if patients == nil {
		return []entities.Patient{}, nil
	}

	return patients, nil
}
