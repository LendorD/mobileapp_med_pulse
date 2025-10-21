package interfaces

import "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

type OneCClient interface {
	GetMedCardByPatientID(patientID string) (*entities.OneCMedicalCard, error)
}
