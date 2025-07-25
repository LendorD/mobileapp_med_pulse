package usecases

import (
	"math"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type EmergencyCallUsecase struct {
	repo interfaces.EmergencyCallRepository
}

func NewEmergencyCallUsecase(repo interfaces.EmergencyCallRepository) interfaces.EmergencyCallUsecase {
	return &EmergencyCallUsecase{repo: repo}
}

// func (u *EmergencyCallUsecase) Create(input models.CreateEmergencyRequest) (entities.EmergencyCall, *errors.AppError) {
// 	emergency := entities.EmergencyCall{
// 		PatientID:       input.PatientID,
// 		Status:          entities.EmergencyStatusScheduled,
// 		Priority:        input.Priority,
// 		Address:         input.Address,
// 		Date:            time.Now(),
// 		Diagnosis:       input.Diagnosis,
// 		Recommendations: input.Recommendations,
// 	}

// 	createdEmergency, err := u.repo.Create(&emergency)
// 	if err != nil {
// 		return entities.EmergencyCall{}, errors.NewDBError("failed to create emergency reception", err)
// 	}

// 	return *createdEmergency, nil
// }

// func (u *EmergencyCallUsecase) AssignDoctor(id, doctorID uint) (entities.EmergencyCall, *errors.AppError) {
// 	emergency, err := u.repo.GetByID(id)
// 	if err != nil {
// 		return entities.EmergencyCall{}, errors.NewDBError("failed to find emergency reception", err)
// 	}

// 	emergency.DoctorID = doctorID
// 	emergency.Status = entities.EmergencyStatusAccepted

// 	updatedEmergency, err := u.repo.Update(emergency)
// 	if err != nil {
// 		return entities.EmergencyCall{}, errors.NewDBError("failed to assign doctor", err)
// 	}

// 	return *updatedEmergency, nil
// }

func (u *EmergencyCallUsecase) GetEmergencyCallsByDoctorAndDate(
	doctorID uint,
	date time.Time,
	page int,
	perPage int,
) (models.FilterResponse[[]models.EmergencyCallShortResponse], error) {
	empty := models.FilterResponse[[]models.EmergencyCallShortResponse]{}
	// Валидация входных параметров
	if doctorID <= 0 {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get doctor",
			errors.ErrEmptyData,
			true,
		)
	}

	if page < 1 {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"page number must be greater than 0",
			errors.ErrDataNotFound,
			true,
		)
	}

	if perPage < 5 {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"Perpage number must be greater than 5",
			errors.ErrDataNotFound,
			true,
		)
	}

	// Получаем данные из репозитория
	calls, total, err := u.repo.GetEmergencyReceptionsByDoctorAndDate(doctorID, date, page, perPage)
	if err != nil {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get receptions",
			errors.ErrEmptyData,
			true,
		)
	}

	// Преобразуем в DTO
	result := make([]models.EmergencyCallShortResponse, len(calls))
	for i, call := range calls {
		result[i] = models.EmergencyCallShortResponse{
			Id:        call.ID,
			CreatedAt: call.CreatedAt.Format(time.RFC3339),
			Phone:     call.Phone,
			Priority:  call.Priority,
			Emergency: call.Emergency,
			Address:   call.Address,
		}
	}

	// Рассчитываем общее количество страниц
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return models.FilterResponse[[]models.EmergencyCallShortResponse]{
		Hits:        result,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalHits:   int(total),
		HitsPerPage: perPage,
	}, nil
}

func (u *EmergencyCallUsecase) UpdateEmergencyCallStatusByID(id uint, newStatus string) (entities.EmergencyCall, error) {
	empty := entities.EmergencyCall{}

	// Обновление статуса в базе
	updateFields := map[string]interface{}{
		"status": newStatus,
	}

	if _, err := u.repo.UpdateEmergencyCall(id, updateFields); err != nil {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to update reception hospital status",
			err,
			true,
		)
	}

	// Получение обновлённой записи
	reception, err := u.repo.GetEmergencyCallByID(id)
	if err != nil {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get updated reception hospital record",
			err,
			true,
		)
	}

	return reception, nil
}

func (u *EmergencyCallUsecase) CloseEmergencyCall(id uint) (entities.EmergencyCall, error) {
	empty := entities.EmergencyCall{}

	// Обновление статуса в базе
	updateFields := map[string]interface{}{
		"priority": nil,
	}

	if _, err := u.repo.UpdateEmergencyCall(id, updateFields); err != nil {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to update reception hospital status",
			err,
			true,
		)
	}

	// Получение обновлённой записи
	reception, err := u.repo.GetEmergencyCallByID(id)
	if err != nil {
		return empty, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get updated reception hospital record",
			err,
			true,
		)
	}

	return reception, nil
}
