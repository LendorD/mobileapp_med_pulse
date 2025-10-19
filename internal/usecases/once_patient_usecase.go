// usecases/onec_patient_list.go

package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	httpClient "github.com/AlexanderMorozov1919/mobileapp/internal/adapters/http/onec"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

type OneCPatientUsecase struct {
	repo       interfaces.PatientRepository
	httpClient httpClient.Client
}

func NewOneCPatientListUsecase(
	repo interfaces.PatientRepository,
	httpClient httpClient.Client,
) interfaces.OneCPatientUsecase {
	return &OneCPatientUsecase{
		repo:       repo,
		httpClient: httpClient,
	}
}

// HandlePatientListUpdate обрабатывает обновление списка пациентов от 1С
func (u *OneCPatientUsecase) HandlePatientListUpdate(ctx context.Context, update entities.PatientListUpdate) error {
	return u.repo.SavePatientList(ctx, update.Patients)
}

// UpdatePatientListFromOneC — отдаёт список пациентов из БД
func (u *OneCPatientUsecase) GetPatientListPage(ctx context.Context, offset, limit int) ([]entities.OneCPatientListItem, error) {
	patients, _, err := u.repo.GetPatientListPage(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

// UpdatePatientListFromOneC — получает список пациентов из 1С и сохраняет в БД
func (u *OneCPatientUsecase) UpdatePatientListFromOneC(ctx context.Context) error {
	req, err := u.httpClient.CreateRequestJSON("GET", "/patients", nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create 1C request: %w", err)
	}

	bodyBytes, _, err := u.httpClient.DoRequest(req)
	if err != nil {
		return fmt.Errorf("failed to send 1C request: %w", err)
	}

	var update entities.PatientListUpdate
	if err := json.Unmarshal(bodyBytes, &update); err != nil {
		return fmt.Errorf("failed to unmarshal 1C response: %w", err)
	}

	if len(update.Patients) == 0 {
		return fmt.Errorf("received empty patient list from 1C: %s", strings.TrimSpace(string(bodyBytes)))
	}

	if err := u.repo.SaveOrUpdatePatientList(ctx, update.Patients); err != nil {
		return fmt.Errorf("failed to save patients: %w", err)
	}

	return nil
}
