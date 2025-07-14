package usecases

import (
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
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

func (u *ReceptionHospitalUsecase) GetPatientsByDoctorID(doc_id uint) ([]entities.Patient, *errors.AppError) {
	if doc_id == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get patient",
			errors.ErrEmptyData,
			true,
		)
	}

	patients, err := u.repo.GetPatientsByDoctorID(doc_id)
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

func (s *ReceptionHospitalUsecase) GetHospitalReceptionsByDoctorAndDate(doctorID uint, date time.Time, page int) ([]models.ReceptionShortResponse, error) {
	// // Валидация номера страницы
	// if page < 1 {
	// 	return nil, errors.New("page must be greater than 0")
	// }

	// // Валидация даты (не раньше текущего дня)
	// if date.Before(time.Now().Truncate(24 * time.Hour)) {
	// 	return nil, errors.New("date cannot be in the past for emergency receptions")
	// }

	// Количество записей на странице
	const perPage = 5 // Можно увеличить для экстренных случаев

	// Получаем данные из репозитория
	receptions, err := s.repo.GetReceptionsHospitalByDoctorAndDate(doctorID, date, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to get emergency receptions: %w", err)
	}

	return receptions, nil
}
