package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type EmergencyReceptionUsecase struct {
	repo interfaces.EmergencyReceptionRepository
}

func NewEmergencyReceptionUsecase(repo interfaces.EmergencyReceptionRepository) interfaces.EmergencyReceptionUsecase {
	return &EmergencyReceptionUsecase{repo: repo}
}

// func (u *EmergencyReceptionUsecase) Create(input models.CreateEmergencyRequest) (entities.EmergencyReception, *errors.AppError) {
// 	emergency := entities.EmergencyReception{
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
// 		return entities.EmergencyReception{}, errors.NewDBError("failed to create emergency reception", err)
// 	}

// 	return *createdEmergency, nil
// }

// func (u *EmergencyReceptionUsecase) AssignDoctor(id, doctorID uint) (entities.EmergencyReception, *errors.AppError) {
// 	emergency, err := u.repo.GetByID(id)
// 	if err != nil {
// 		return entities.EmergencyReception{}, errors.NewDBError("failed to find emergency reception", err)
// 	}

// 	emergency.DoctorID = doctorID
// 	emergency.Status = entities.EmergencyStatusAccepted

// 	updatedEmergency, err := u.repo.Update(emergency)
// 	if err != nil {
// 		return entities.EmergencyReception{}, errors.NewDBError("failed to assign doctor", err)
// 	}

// 	return *updatedEmergency, nil
// }
