package usecases

import (
	"errors"

	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/interfaces"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DoctorUsecase struct {
	repo interfaces.DoctorRepository
}

func NewDoctorUsecase(repo interfaces.DoctorRepository) interfaces.DoctorUsecase {
	return &DoctorUsecase{repo: repo}
}

func (u *DoctorUsecase) Create(input models.CreateDoctorRequest) (entities.Doctor, *errors.AppError) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.Doctor{}, errors.NewInternalServerError("password hashing failed", err)
	}

	doctor := entities.Doctor{
		FullName:       input.FullName,
		Login:          input.Login,
		Email:          input.Email,
		PasswordHash:   string(passwordHash),
		Specialization: input.Specialization,
	}

	createdDoctor, err := u.repo.Create(&doctor)
	if err != nil {
		return entities.Doctor{}, errors.NewDBError("failed to create doctor", err)
	}

	return *createdDoctor, nil
}

func (u *DoctorUsecase) GetByID(id uint) (entities.Doctor, *errors.AppError) {
	doctor, err := u.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Doctor{}, errors.NewNotFoundError("doctor not found")
		}
		return entities.Doctor{}, errors.NewDBError("failed to get doctor", err)
	}
	return *doctor, nil
}

func (u *DoctorUsecase) Update(input models.UpdateDoctorRequest) (entities.Doctor, *errors.AppError) {
	doctor, err := u.repo.GetByID(input.ID)
	if err != nil {
		return entities.Doctor{}, errors.NewDBError("failed to find doctor", err)
	}

	if input.FullName != "" {
		doctor.FullName = input.FullName
	}
	if input.Email != "" {
		doctor.Email = input.Email
	}
	if input.Specialization != "" {
		doctor.Specialization = input.Specialization
	}

	updatedDoctor, err := u.repo.Update(doctor)
	if err != nil {
		return entities.Doctor{}, errors.NewDBError("failed to update doctor", err)
	}

	return *updatedDoctor, nil
}

func (u *DoctorUsecase) Delete(id uint) *errors.AppError {
	if err := u.repo.Delete(id); err != nil {
		return errors.NewDBError("failed to delete doctor", err)
	}
	return nil
}

func (u *DoctorUsecase) GetFiltered(page, count int, order, filter string) (models.FilterResponse[[]entities.Doctor], *errors.AppError) {
	// Реализация фильтрации аналогична примеру BookingGroupUsecase
	// ...
}
