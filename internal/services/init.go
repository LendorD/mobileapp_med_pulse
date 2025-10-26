package services

import (
	"log"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type Service struct {
	interfaces.ParamsParserService
	interfaces.FilterBuilderService
	interfaces.ImageService
}

func NewService(cfg *config.Config) interfaces.Service {
	parser := NewParamsParser()
	imageSvc, err := NewImageService(cfg.MinIO)
	if err != nil {
		log.Fatalf("Failed to initialize ImageService: %v", err)
	}

	return Service{
		ParamsParserService:  parser,
		FilterBuilderService: NewFilterBuilder(parser),
		ImageService:         imageSvc,
	}
}
