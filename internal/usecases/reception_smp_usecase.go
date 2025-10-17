package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type ReceptionSmpUsecase struct {
	recepSmpRepo interfaces.ReceptionSmpRepository
	patientRepo  interfaces.PatientRepository
	doctorRepo   interfaces.DoctorRepository
}

func NewReceptionSmpUsecase(recepRepo interfaces.ReceptionSmpRepository, patientRepo interfaces.PatientRepository) interfaces.ReceptionSmpUsecase {
	return &ReceptionSmpUsecase{
		recepSmpRepo: recepRepo,
		patientRepo:  patientRepo,
	}
}
