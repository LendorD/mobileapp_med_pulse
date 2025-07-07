package usecases

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	_ "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	_ "github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type UseCases struct {
}

func NewUsecases(r interfaces.Repository, s interfaces.Service, conf *config.Config) interfaces.Usecases {

	// Создание структуры UseCases
	return &UseCases{}

}
