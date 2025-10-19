package interfaces

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type Usecases interface {
	ReceptionSmpUsecase
	MedCardUsecase
	AuthUsecase
	OneCWebhookUsecase
	OneCPatientUsecase
}

type OneCPatientUsecase interface {
	HandlePatientListUpdate(ctx context.Context, update entities.PatientListUpdate) error
	GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, error)
}

type OneCWebhookUsecase interface {
	HandleReceptionsUpdate(ctx context.Context, update models.Call) error
	GetInterestedUserIDs(callID int) []uint
}

type ReceptionSmpUsecase interface {
}

type MedCardUsecase interface {
	GetMedCardByPatientID(ctx context.Context, patientID string) (*entities.OneCMedicalCard, error)
	UpdateMedicalCard(ctx context.Context, card *entities.OneCMedicalCard) error
}

type AuthUsecase interface {
	SyncUsers(ctx context.Context, users []entities.AuthUser) error
	LoginDoctor(ctx context.Context, phone, password string) (*models.DoctorAuthResponse, *errors.AppError)
}
