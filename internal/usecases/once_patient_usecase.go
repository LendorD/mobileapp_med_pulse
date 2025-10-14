// usecases/onec_patient_list.go

package usecases

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type OneCPatientUsecase struct {
	cacheRepo interfaces.OneCCacheRepository
}

func NewOneCPatientListUsecase(
	cacheRepo interfaces.OneCCacheRepository,
) interfaces.OneCPatientUsecase {
	return &OneCPatientUsecase{
		cacheRepo: cacheRepo,
	}
}

// HandlePatientListUpdate обрабатывает обновление списка пациентов от 1С
func (u *OneCPatientUsecase) HandlePatientListUpdate(ctx context.Context, update models.PatientListUpdate) error {
	// Сохраняем список пациентов в кеш
	return u.cacheRepo.SavePatientList(ctx, update.Patients)
}

func (u *OneCPatientUsecase) GetPatientList(ctx context.Context) ([]models.PatientListItem, error) {
	return u.cacheRepo.GetPatientList(ctx)
}
