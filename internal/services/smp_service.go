package services

import (
	"errors"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/repository"
)

type ReceptionResponce struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	IsMale    bool      `json:"is_male"`
	Address   string    `json:"address"`
	Date      time.Time `json:"date"`
}

type smpService struct {
	smpRepo     repository.ReceptionRepository
	patientRepo repository.PatientRepository
}

func NewSmpService(smpRepo repository.ReceptionRepository) SmpService {
	return &smpService{smpRepo: smpRepo}
}

func (s *smpService) GetCallings(doctorID uint) ([]ReceptionResponce, error) {
	receptions, err := s.smpRepo.GetAllByDoctorID(doctorID)
	if err != nil {
		return nil, errors.New("failed to get receptions")
	}

	var response []ReceptionResponce

	for _, reception := range receptions {
		patient, err := s.patientRepo.GetByID(reception.PatientID)
		if err != nil {
			return nil, errors.New("failed to get patient data")
		}

		respItem := ReceptionResponce{
			ID:        reception.ID,
			FullName:  patient.FullName,
			BirthDate: patient.BirthDate,
			IsMale:    patient.IsMale,
			Address:   reception.Address,
			Date:      reception.Date,
		}

		response = append(response, respItem)
	}

	return response, nil
}
