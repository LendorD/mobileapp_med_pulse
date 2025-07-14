package usecases

import (
	"log"
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

func (u *ReceptionHospitalUsecase) GetHospitalReceptionsByDoctorAndDate(doctorID uint, date time.Time, page int) ([]models.ReceptionShortResponse, error) {
	const perPage = 5

	// Валидация входных параметров
	if doctorID == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get patient",
			errors.ErrEmptyData,
			true,
		)
	}

	if page < 1 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"page number must be greater than 0",
			errors.ErrDataNotFound,
			true,
		)
	}

	// Получаем данные из репозитория
	receptions, err := u.repo.GetReceptionsHospitalByDoctorAndDate(doctorID, date, page, perPage)
	if err != nil {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get receptions",
			errors.ErrEmptyData,
			true,
		)
	}

	result := make([]models.ReceptionShortResponse, len(receptions))
	for i, reception := range receptions {
		// Получаем имя пациента (с проверкой что Patient загружен)
		patientName := ""
		if reception.Patient.ID != 0 { // Проверка что связь загружена
			patientName = reception.Patient.FullName
		}

		// Получаем статус в читаемом формате
		statusText := getStatusText(reception.Status)

		// Форматируем дату
		formattedDate := reception.Date.Format("02.01.2006 15:04") // Формат "15.10.2023 14:30"

		result[i] = models.ReceptionShortResponse{
			Date:        formattedDate,
			Status:      statusText,
			PatientName: patientName,
		}
	}

	return result, nil
}

func getStatusText(status entities.ReceptionStatus) string {
	switch status {
	case entities.StatusScheduled:
		return "Запланирован"
	case entities.StatusCompleted:
		return "Завершен"
	case entities.StatusCancelled:
		return "Отменен"
	case entities.StatusNoShow:
		return "Не явился"
	default:
		return string(status)
	}
}
