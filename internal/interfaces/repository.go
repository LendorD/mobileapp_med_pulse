package interfaces

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"gorm.io/gorm"
)

type Repository interface {
	AuthRepository
	DoctorRepository
	PatientRepository
	ReceptionSmpRepository
	MedicalCardRepository
	TxManager
	FileRepository
	FlgRepository
	AnalysisRepository
}

type AnalysisRepository interface {
	GetAnalysisByID(ctx context.Context, id uint) (*entities.Analysis, error)
	GetAllAnalysisIDs(ctx context.Context) ([]uint, error)
	GetAllAnalyses(ctx context.Context) ([]entities.Analysis, error)

	UpdateAnalysisOrder(ctx context.Context, order *entities.AnalysisOrder) error
	CreateAnalysisOrder(ctx context.Context, order *entities.AnalysisOrder) error
	CreateAnalysisItems(ctx context.Context, items []entities.AnalysisOrderItem) error

	GetAnalysisOrderByID(ctx context.Context, id uint) (*entities.AnalysisOrder, error)
	GetOrderItemsByOrderID(ctx context.Context, orderID uint) ([]entities.AnalysisOrderItem, error)
	UpsertOrderItems(ctx context.Context, items []entities.AnalysisOrderItem) error
}

type FlgRepository interface {
	CreateFlg(ctx context.Context, flg *entities.Flg) error
	GetFlgByPatientID(ctx context.Context, patientID uint) ([]entities.Flg, error)
	GetFlgByID(ctx context.Context, id uint) (*entities.Flg, error)
	Delete(ctx context.Context, id uint) error
}

type TxManager interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	GetTransaction(ctx context.Context) *gorm.DB
}

type FileRepository interface {
	CreateFile(ctx context.Context, file *entities.File) error
	GetFileByID(ctx context.Context, id uint) (*entities.File, error)
	DeleteFile(ctx context.Context, id uint) error
}

type MedicalCardRepository interface {
	SaveMedicalCard(ctx context.Context, card *entities.OneCMedicalCard) error
	GetMedicalCard(ctx context.Context, patientID string) (*entities.OneCMedicalCard, error)
	DeleteMedicalCard(ctx context.Context, patientID string) error
}

// updated to match the new structure
type DoctorRepository interface {
	GetDoctorByID(ctx context.Context, id uint) (entities.Doctor, error)
	GetDoctorByLogin(ctx context.Context, login string) (entities.Doctor, error)
}

// updated to match the new structure
type ReceptionSmpRepository interface {
	// Вызовы (скорая)
	GetUndeliveredReceptions(ctx context.Context) ([]entities.OneCReception, error)
	UpdateStatus(ctx context.Context, callID, status string) error
	SaveReceptions(ctx context.Context, callID string, reception entities.OneCReception) error
	GetReceptions(ctx context.Context, callID string) ([]models.Patient, error)
}

// updated to match the new structured
type PatientRepository interface {
	// Список пациентов
	SavePatientList(ctx context.Context, patients []entities.OneCPatientListItem) error
	SaveOrUpdatePatientList(ctx context.Context, patients []entities.OneCPatientListItem) error
	GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, int64, error)
}

type AuthRepository interface {
	SaveUsers(ctx context.Context, users []entities.AuthUser) error
	GetUserByLogin(ctx context.Context, login string) (*entities.AuthUser, error)
}
