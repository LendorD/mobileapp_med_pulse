package services

import "github.com/AlexanderMorozov1919/mobileapp/internal/repository"

type smpService struct {
	smpRepo repository.ReceptionRepository
}

func NewSmpService(smpRepo repository.ReceptionRepository) SmpService {
	return &smpService{smpRepo: smpRepo}
}
