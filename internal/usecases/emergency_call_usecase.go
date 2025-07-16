package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
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

func (s *EmergencyCallUsecase) GetEmergencyCallsByDoctorAndDate(doctorID uint, date time.Time, page int) ([]models.EmergencyCallShortResponse, error) {
	// Валидация номера страницы
	if page < 1 {
		return nil, errors.New("page must be greater than 0")
	}

	// Валидация даты (не раньше текущего дня)
	if date.Before(time.Now().Truncate(24 * time.Hour)) {
		return nil, errors.New("date cannot be in the past for emergency receptions")
	}

	// Количество записей на странице
	const perPage = 5 // Можно увеличить для экстренных случаев

	// Получаем данные из репозитория
	receptions, err := s.repo.GetEmergencyCallsByDoctorAndDate(doctorID, date, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to get emergency receptions: %w", err)
	}

	return receptions, nil
}
