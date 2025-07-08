package interfaces

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"

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

type AllergyUsecase interface {
	AddAllergyToPatient(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError)
	GetAllergyByPatientID(patientID uint) ([]entities.Allergy, *errors.AppError)
	RemoveAllergyFromPatient(patientID, allergyID uint) *errors.AppError
	UpdateAllergyDescription(patientID, allergyID uint, description string) (entities.PatientsAllergy, *errors.AppError)
}

type ContactInfoUsecase interface{}

type DoctorUsecase interface {
	CreateDoctor(doctor models.CreateDoctorRequest) (entities.Doctor, *errors.AppError)
	GetDoctorByID(doctorId uint) (entities.Doctor, *errors.AppError)
	UpdateDoctor(doctor models.UpdateDoctorRequest) (entities.Doctor, *errors.AppError)
	DeleteDoctor(doctorId uint) *errors.AppError
}

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

type PatientUsecase interface {
	CreatPatient(input *models.CreatePatientRequest) (entities.Patient, *errors.AppError)
	GetPatientByID(id uint) (entities.Patient, *errors.AppError)
	UpdatePatient(input *models.UpdatePatientRequest) (entities.Patient, *errors.AppError)
	DeletePatient(id uint) *errors.AppError
}

type PersonalInfoUsecase interface{}

type ReceptionUsecase interface{}
