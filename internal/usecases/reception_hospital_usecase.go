package usecases

import (
	"log"

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
		log.Println(reception.Doctor)
		log.Println(reception.Patient.FullName)
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
