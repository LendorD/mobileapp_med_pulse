package models

import "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

type PatientListResponse struct {
	Patient []entities.OneCPatientListItem
}
