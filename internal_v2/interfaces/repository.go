package interfaces

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
	"time"
	//"github.com/AlexanderMorozov1919/mobileapp/internal/models"
)

type DoctorRepository interface {
	Create(doctor *entities.Doctor) error
	Update(docor *entities.Doctor) error
	Delete(id uint) error
	GetByID(id uint) (*entities.Doctor, error)
	GetName(id uint) (string, error)
	GetByLogin(login string) (*entities.Doctor, error)
	GetSpecialization(id uint) (string, error)
	GetPassHash(id uint) (string, error)
}

type PersonalInfoRepository interface {
	Create(info *entities.PersonalInfo) error
	Update(info *entities.PersonalInfo) error
	Delete(id uint) error

	GetByID(id uint) (*entities.PersonalInfo, error)
	GetByPatientID(patientID uint) (*entities.PersonalInfo, error)
}

type EmergencyReceptionRepository interface {
	Create(er *entities.EmergencyReception) error
	Update(er *entities.EmergencyReception) error
	Delete(id uint) error

	GetByID(id uint) (*entities.EmergencyReception, error)
	GetByDoctorID(doctorID uint) ([]entities.EmergencyReception, error)
	GetByPatientID(patientID uint) ([]entities.EmergencyReception, error)
	GetByDateRange(start, end time.Time) ([]entities.EmergencyReception, error)
	GetPriorityCases() ([]entities.EmergencyReception, error)
}

type MedServiceRepository interface {
	Create(service *entities.MedService) error
	Update(service *entities.MedService) error
	Delete(id uint) error

	GetByID(id uint) (*entities.MedService, error)
	GetByName(name string) (*entities.MedService, error)
	GetAll() ([]entities.MedService, error)
}

type EmergencyReceptionMedServicesRepository interface {
	Create(link *entities.EmergencyReceptionMedServices) error
	Delete(id uint) error

	GetByEmergencyReceptionID(erID uint) ([]entities.EmergencyReceptionMedServices, error)
}

type ReceptionRepository interface {
	Create(reception *entities.Reception) error
	Update(reception *entities.Reception) error
	Delete(id uint) error

	GetByID(id uint) (*entities.Reception, error)
	GetByDoctorID(doctorID uint) ([]entities.Reception, error)
	GetByPatientID(patientID uint) ([]entities.Reception, error)
	GetByDateRange(start, end time.Time) ([]entities.Reception, error)
}

type PatientRepository interface {
	Create(patient *entities.Patient) error
	Update(patient *entities.Patient) error
	Delete(id uint) error

	GetByID(id uint) (*entities.Patient, error)
	GetAll() ([]entities.Patient, error)
	GetByFullName(name string) ([]entities.Patient, error)
}

type ContactInfoRepository interface {
	Create(info *entities.ContactInfo) error
	Update(info *entities.ContactInfo) error
	Delete(id uint) error

	GetByID(id uint) (*entities.ContactInfo, error)
	GetByPatientID(patientID uint) (*entities.ContactInfo, error)
}

type AllergyRepository interface {
	Create(allergy *entities.Allergy) error
	Update(allergy *entities.Allergy) error
	Delete(id uint) error

	GetByID(id uint) (*entities.Allergy, error)
	GetByName(name string) (*entities.Allergy, error)
	GetAll() ([]entities.Allergy, error)
}

type PatientsAllergyRepository interface {
	Create(pa *entities.PatientsAllergy) error
	Update(pa *entities.PatientsAllergy) error
	Delete(id uint) error

	GetByID(id uint) (*entities.PatientsAllergy, error)
	GetByPatientID(patientID uint) ([]entities.PatientsAllergy, error)
}
