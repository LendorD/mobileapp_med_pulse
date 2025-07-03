package services

import (
	"log"

	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
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

func (s *patientService) GetAllPatientsByDoctorID(doctorID uint) ([]models.ShortPatientResponse, error) {
	s.logger.Println("[INFO] Получение всех пользователей")
	patients, err := s.patientRepository.GetAllPatientsByDoctorID(doctorID)
	if err != nil {
		s.logger.Printf("[ERROR] Ошибка при получении пользователей: %v", err)
		return nil, err
	}

	var response = make([]models.ShortPatientResponse, len(patients))
	for _, patient := range patients {
		response = append(response, models.ShortPatientResponse{
			ID:        patient.ID,
			FullName:  patient.FullName,
			BirthDate: patient.BirthDate,
			IsMale:    patient.IsMale,
		})
	}
	return response, nil
}
