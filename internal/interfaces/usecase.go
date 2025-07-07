package interfaces

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

type MedServiceUsecase interface{}

type PatientUsecase interface{}

type PersonalInfoUsecase interface{}

type ReceptionUsecase interface{}
