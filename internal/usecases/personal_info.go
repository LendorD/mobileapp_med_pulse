package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type PersonalInfoUsecase struct {
	repo interfaces.PersonalInfoRepository
}

func NewPersonalInfoUsecase(repo interfaces.PersonalInfoRepository) interfaces.PersonalInfoUsecase {
	return &PersonalInfoUsecase{repo: repo}
}

// func (u *PersonalInfoUsecase) GetByPatientID(patientID uint) (entities.PersonalInfo, *errors.AppError) {
// 	info, err := u.repo.GetByPatientID(patientID)
// 	if err != nil {
// 		if erro.Is(err, gorm.ErrRecordNotFound) {
// 			return entities.PersonalInfo{}, errors.NewNotFoundError("personal info not found")
// 		}
// 		return entities.PersonalInfo{}, errors.NewDBError("failed to get personal info", err)
// 	}
// 	return *info, nil
// }

// func (u *PersonalInfoUsecase) Update(input models.UpdatePersonalInfoRequest) (entities.PersonalInfo, *errors.AppError) {
// 	info, err := u.repo.GetByPatientID(input.PatientID)
// 	if err != nil {
// 		return entities.PersonalInfo{}, errors.NewDBError("failed to find personal info", err)
// 	}

// 	info.PassportSeries = input.PassportSeries
// 	errorsinfo.SNILS = input.SNILS
// 	info.OMS = input.OMS

// 	updatedInfo, err := u.repo.Update(info)
// 	if err != nil {
// 		return entities.PersonalInfo{}, errors.NewDBError("failed to update personal info", err)
// 	}

// 	return *updatedInfo, nil
// }
