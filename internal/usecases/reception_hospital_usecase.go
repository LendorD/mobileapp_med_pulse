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

// []models.ReceptionHospitalResponse
func (u *ReceptionHospitalUsecase) GetReceptionsHospitalByPatientID(patientId uint) ([]models.ReceptionHospitalResponse, *errors.AppError) {
	if patientId == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get patient",
			errors.ErrEmptyData,
			true,
		)
	}

	receptions, err := u.repo.GetReceptionHospitalByPatientID(patientId)
	// receptionsSMP, err := u.repo.GetReceptionHospitalByPatientID(patientId)
	if err != nil {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get receptions",
			errors.ErrEmptyData,
			true,
		)
	}
	var responses []models.ReceptionHospitalResponse
	for _, reception := range receptions {
		fmt.Printf("Doctor: %+v\n", reception.Doctor)
		fmt.Printf("Patient: %+v\n", reception.Patient)
		response := models.ReceptionHospitalResponse{
			Doctor: models.DoctorInfoResponse{
				FullName:       reception.Doctor.FullName,
				Specialization: reception.Doctor.Specialization,
			},
			Patient: models.ShortPatientResponse{
				ID:        reception.Patient.ID,
				FullName:  reception.Patient.FullName,
				BirthDate: reception.Patient.BirthDate,
				IsMale:    reception.Patient.IsMale,
			},
			Diagnosis:       reception.Diagnosis,
			Recommendations: reception.Recommendations,
			Date:            reception.Date,
		}
		responses = append(responses, response)
	}

	if len(responses) == 0 {
		return []models.ReceptionHospitalResponse{}, nil
	}

	return responses, nil
}

// func (u *ReceptionHospitalUsecase) UpdateReceptionHospital(input *models.UpdateReceptionHospitalRequest) (models.ReceptionHospitalResponse, *errors.AppError) {
// 	recepHospUpdate := map[string]interface{}{
// 		"id":              input.ID,
// 		"diagnosis":       input.Diagnosis,
// 		"recommendations": input.Recommendations,
// 		"status":          input.Status,
// 	}

// 	if _, err := u.repo.UpdateReceptionHospital(input.ID, recepHospUpdate); err != nil {
// 		return models.ReceptionHospitalResponse{}, errors.NewAppError(
// 			errors.InternalServerErrorCode,
// 			"failed to update reception hospital data",
// 			err,
// 			true,
// 		)
// 	}

// 	recepHospResponse, err := u.repo.GetReceptionHospitalByID(input.ID)
// 	if err != nil {
// 		return models.ReceptionHospitalResponse{}, errors.NewAppError(
// 			errors.InternalServerErrorCode,
// 			"failed to get updated reception hospital data",
// 			err,
// 			true,
// 		)
// 	}
// 	return models.ReceptionHospitalResponse{
// 		Doctor: models.DoctorInfoResponse{
// 			FullName:       recepHospResponse.Doctor.FullName,
// 			Specialization: recepHospResponse.Doctor.Specialization,
// 		},
// 		Patient: models.ShortPatientResponse{
// 			ID:        recepHospResponse.Patient.ID,
// 			FullName:  recepHospResponse.Patient.FullName,
// 			BirthDate: recepHospResponse.Patient.BirthDate,
// 			IsMale:    recepHospResponse.Patient.IsMale,
// 		},
// 		Diagnosis:       recepHospResponse.Diagnosis,
// 		Recommendations: recepHospResponse.Recommendations,
// 		Date:            recepHospResponse.Date,
// 	}, nil
// }

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
