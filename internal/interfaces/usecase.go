package interfaces

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"

	"time"
)

type Usecases interface {
	AllergyUsecase
	ContactInfoUsecase
	DoctorUsecase
	EmergencyReceptionUsecase
	EmergencyReceptionMedServicesUsecase
	MedServiceUsecase
	PatientUsecase
	PersonalInfoUsecase
	ReceptionUsecase
}

type AllergyUsecase interface{}

type ContactInfoUsecase interface{}

type DoctorUsecase interface{}

type EmergencyReceptionUsecase interface{}

type EmergencyReceptionMedServicesUsecase interface{}

// ReceptionService определяет контракт для работы с записями на прием
type ReceptionService interface {
	CreateReception(reception *models.Reception) error
	UpdateReception(reception *models.Reception) error
	CancelReception(id uint, reason string) error
	CompleteReception(id uint, diagnosis string, recommendations string) error
	MarkAsNoShow(id uint) error
	GetReceptionByID(id uint) (*models.Reception, error)
	GetDoctorReceptions(doctorID uint, date *time.Time) ([]models.Reception, error)
	GetPatientReceptions(patientID uint) ([]models.Reception, error)
	GetReceptionsByStatus(status entities.ReceptionStatus) ([]models.Reception, error)
	GetReceptionsByDoctorAndDate(doctorID uint, date time.Time) ([]models.Reception, error)
}
type MedServiceUsecase interface{}

type PatientUsecase interface{}

type PersonalInfoUsecase interface{}

type ReceptionUsecase interface{}
