package services

import "github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type Service struct {
	/*
		interfaces.TimeValidatorService
		interfaces.ParseFilterService

	*/
}

func NewService(parentLogger *logging.Logger) interfaces.Service {
	//logger := logging.NewModuleLogger("SERVICES", "GENERAL", parentLogger)

	return Service{}
}
