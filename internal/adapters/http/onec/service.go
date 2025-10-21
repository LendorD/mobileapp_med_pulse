package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
)

func (c *OneCClient) GetMedCardByPatientID(patientID string) (*entities.OneCMedicalCard, error) {
	endpoint := fmt.Sprintf("/medical-card/%s", patientID)
	req, err := c.CreateRequestJSON(http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create 1C request: %w", err)
	}

	body, _, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("1C request error: %w", err)
	}

	var patientCard entities.OneCMedicalCard
	if err := json.Unmarshal(body, &patientCard); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &patientCard, nil
}

func (c *OneCClient) UpdateMedCardByPatientID(patientID string, card *entities.OneCMedicalCard) error {
	body, err := json.Marshal(card)

	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	endpoint := fmt.Sprintf("/medical-card/%s", patientID)

	req, err := c.CreateRequestJSON(http.MethodPost, endpoint, nil, nil, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("1C update error: %w", err)
	}

	_, _, err = c.DoRequest(req)
	if err != nil {
		return fmt.Errorf("1C request error: %w", err)
	}

	return nil
}
