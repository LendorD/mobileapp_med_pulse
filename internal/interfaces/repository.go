package interfaces

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
)

type Repository interface {
	AuthRepository
	DoctorRepository
	PatientRepository
	ReceptionSmpRepository
	MedicalCardRepository
}

type MedicalCardRepository interface {
	SaveMedicalCard(ctx context.Context, card *entities.OneCMedicalCard) error
	GetMedicalCard(ctx context.Context, patientID string) (*entities.OneCMedicalCard, error)
	DeleteMedicalCard(ctx context.Context, patientID string) error
}

// updated to match the new structure
type DoctorRepository interface {
	GetDoctorByID(id uint) (entities.Doctor, error)
	GetDoctorByLogin(login string) (entities.Doctor, error)
}

// updated to match the new structure
type ReceptionSmpRepository interface {
	// Вызовы (скорая)
	SaveReceptions(ctx context.Context, callID string, receptions []models.Patient) error
	GetReceptions(ctx context.Context, callID string) ([]models.Patient, error)
}

// updated to match the new structured
type PatientRepository interface {
	// Список пациентов
	SavePatientList(ctx context.Context, patients []entities.OneCPatientListItem) error
	GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, int64, error)
}

type AuthRepository interface {
	SaveUsers(ctx context.Context, users []entities.AuthUser) error
	GetUserByLogin(ctx context.Context, login string) (*entities.AuthUser, error)
}
