package services

import (
	"log"

	"github.com/AlexanderMorozov1919/mobileapp/internal/repository"
)

type patientService struct {
	patientRepository repository.PatientRepository
	logger            *log.Logger
}

func NewPatientService(patientRepository repository.PatientRepository, logger *log.Logger) PatientService {
	return &patientService{
		patientRepository: patientRepository,
		logger:            logger,
	}
}
