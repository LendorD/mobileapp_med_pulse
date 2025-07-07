package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type ContactInfoUsecase struct {
	repo interfaces.ContactInfoRepository
}

func NewContactInfoUsecase(repo interfaces.ContactInfoRepository) interfaces.ContactInfoUsecase {
	return &ContactInfoUsecase{repo: repo}
}

// func (u *ContactInfoUsecase) GetByPatientID(patientID uint) (entities.ContactInfo, *errors.AppError) {
// 	info, err := u.repo.GetByPatientID(patientID)
// 	if err != nil {
// 		if error.Is(err, gorm.ErrRecordNotFound) {
// 			return entities.ContactInfo{}, errors.NewNotFoundError("contact info not found")
// 		}
// 		return entities.ContactInfo{}, errors.NewDBError("failed to get contact info", err)
// 	}
// 	return *info, nil
// }

// func (u *ContactInfoUsecase) Update(input models.UpdateContactInfoRequest) (entities.ContactInfo, *errors.AppError) {
// 	info, err := u.repo.GetByPatientID(input.PatientID)
// 	if err != nil {
// 		return entities.ContactInfo{}, errors.NewDBError("failed to find contact info", err)
// 	}

// 	info.Phone = input.Phone
// 	errorsinfo.Email = input.Email
// 	info.Address = input.Address

// 	updatedInfo, err := u.repo.Update(info)
// 	if err != nil {
// 		return entities.ContactInfo{}, errors.NewDBError("failed to update contact info", err)
// 	}

// 	return *updatedInfo, nil
// }
