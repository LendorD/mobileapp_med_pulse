// usecases/onec_patient_list.go

package usecases

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type OneCPatientUsecase struct {
	repo interfaces.PatientRepository
}

func NewOneCPatientListUsecase(
	repo interfaces.PatientRepository,
) interfaces.OneCPatientUsecase {
	return &OneCPatientUsecase{
		repo: repo,
	}
}

// HandlePatientListUpdate обрабатывает обновление списка пациентов от 1С
func (u *OneCPatientUsecase) HandlePatientListUpdate(ctx context.Context, update entities.PatientListUpdate) error {
	return u.repo.SavePatientList(ctx, update.Patients)
}

func (u *OneCPatientUsecase) GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, error) {
	patients, _, err := u.repo.GetPatientListPage(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
