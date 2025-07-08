package interfaces

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

type Repository interface {
	AllergyRepository
	DoctorRepository
	MedServiceRepository
	EmergencyReceptionMedServicesRepository
	PatientRepository
	PatientsAllergyRepository
	ContactInfoRepository
	EmergencyReceptionRepository
	PersonalInfoRepository
	ReceptionRepository
}

type DoctorRepository interface {
	CreateDoctor(doctor *entities.Doctor) (uint, *errors.AppError)
	UpdateDoctor(id uint, updateMap map[string]interface{}) (uint, *errors.AppError)
	DeleteDoctor(id uint) *errors.AppError
	GetDoctorByID(id uint) (*entities.Doctor, *errors.AppError)
	GetDoctorName(id uint) (string, *errors.AppError)
	GetDoctorByLogin(login string) (*entities.Doctor, *errors.AppError)
	GetDoctorSpecialization(id uint) (string, *errors.AppError)
	GetDoctorPassHash(id uint) (string, *errors.AppError)
}

type PersonalInfoRepository interface {
	CreatePersonalInfo(info *entities.PersonalInfo) error
	UpdatePersonalInfo(info *entities.PersonalInfo) error
	DeletePersonalInfo(id uint) error

	GetPersonalInfoByID(id uint) (*entities.PersonalInfo, error)
	GetPersonalInfoByPatientID(patientID uint) (*entities.PersonalInfo, error)
}

type EmergencyReceptionRepository interface {
	CreateEmergencyReception(er *entities.EmergencyReception) error
	UpdateEmergencyReception(er *entities.EmergencyReception) error
	DeleteEmergencyReception(id uint) error

	GetEmergencyReceptionByID(id uint) (*entities.EmergencyReception, error)
	GetEmergencyReceptionByDoctorID(doctorID uint) ([]entities.EmergencyReception, error)
	GetEmergencyReceptionByPatientID(patientID uint) ([]entities.EmergencyReception, error)
	GetEmergencyReceptionByDateRange(start, end time.Time) ([]entities.EmergencyReception, error)
	GetEmergencyReceptionPriorityCases() ([]entities.EmergencyReception, error)
}

type MedServiceRepository interface {
	CreateMedService(service *entities.MedService) error
	UpdateMedService(service *entities.MedService) error
	DeleteMedService(id uint) error

	GetMedServiceByID(id uint) (*entities.MedService, error)
	GetMedServiceByName(name string) (*entities.MedService, error)
	GetAllMedService() ([]entities.MedService, error)
}

type EmergencyReceptionMedServicesRepository interface {
	CreateEmergencyReceptionMedServices(link *entities.EmergencyReceptionMedServices) error
	DeleteEmergencyReceptionMedServices(id uint) error

	GetEmergencyReceptionMedServicesByEmergencyReceptionID(erID uint) ([]entities.EmergencyReceptionMedServices, error)
}

type ReceptionRepository interface {
	CreateReception(reception *entities.Reception) error
	UpdateReception(reception *entities.Reception) error
	DeleteReception(id uint) error

	GetReceptionByID(id uint) (*entities.Reception, error)
	GetReceptionByDoctorID(doctorID uint) ([]entities.Reception, error)
	GetReceptionByPatientID(patientID uint) ([]entities.Reception, error)
	GetReceptionByDateRange(start, end time.Time) ([]entities.Reception, error)
}

type PatientRepository interface {
	CreatePatient(patient entities.Patient) (uint, *errors.AppError)
	UpdatePatient(id uint, updateMap map[string]interface{}) (uint, *errors.AppError)
	DeletePatient(id uint) *errors.AppError
	GetPatientByID(id uint) (entities.Patient, *errors.AppError)
	//GetAllPatients() ([]entities.Patient, *errors.AppError)
	//GetPatientsByFullName(name string) ([]entities.Patient, *errors.AppError)
}

type ContactInfoRepository interface {
	CreateContactInfo(info *entities.ContactInfo) error
	UpdateContactInfo(info *entities.ContactInfo) error
	DeleteContactInfo(id uint) error

	GetContactInfoByID(id uint) (*entities.ContactInfo, error)
	GetContactInfoByPatientID(patientID uint) (*entities.ContactInfo, error)
}

type AllergyRepository interface {
	CreateAllergy(allergy *entities.Allergy) error
	UpdateAllergy(allergy *entities.Allergy) error
	DeleteAllergy(id uint) error

	GetPatientAllergiesByID(id uint) ([]entities.Allergy, error)
	//GetPatientAllergyByID(id uint) (*entities.Allergy, error)
	GetAllergyByName(name string) (*entities.Allergy, error)
	GetAllAllergy() ([]entities.Allergy, error)
	//GetPatientAllergyByID(id uint) (*entities.PatientsAllergy, error)
}

type PatientsAllergyRepository interface {
	CreatePatientsAllergy(pa *entities.PatientsAllergy) error
	UpdatePatientsAllergy(pa *entities.PatientsAllergy) error
	DeletePatientsAllergy(id uint) error
	ExistsAllergy(patientID, allergyID uint) (bool, error)
	//GetPatientsAllergyByAllergyID(id uint) (*entities.PatientsAllergy, error)
	//GetPatientsAllergiesByPatientID(patientID uint) ([]entities.PatientsAllergy, error)
	//GetAllergyByPatientID(patientID uint) ([]entities.Allergy, error)
}
