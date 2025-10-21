package httpClient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	httpUrl "net/url"
	"strings"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

// Client — HTTP-клиент для взаимодействия с 1С
type OneCClient struct {
	Host       string
	HTTPClient *http.Client
	username   string
	password   string
}

// NewClient создаёт новый клиент для 1С
func NewOneCClient(cfg config.OneCConfig) interfaces.OneCClient {
	// Создаём HTTP-клиент с таймаутом
	client := &http.Client{
		Timeout: cfg.Timeout,
	}

	return &OneCClient{
		Host:       cfg.BaseURL,
		HTTPClient: client,
		username:   cfg.Username,
		password:   cfg.Password,
	}
}

// createRequestJSON to create request with json
func (s *OneCClient) CreateRequestJSON(httpMethod, endpoint string, queryParams, additionalHeaders map[string]string, reqBody io.Reader) (*http.Request, error) {
	ctx := context.Background()

	proccesedEndpoint, _ := strings.CutPrefix(endpoint, "/")

	url := fmt.Sprintf("%s/%s", s.Host, proccesedEndpoint)

	if len(queryParams) > 0 {
		params := httpUrl.Values{}
		for param, value := range queryParams {
			params.Add(param, value)
		}
		url += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range additionalHeaders {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (s *OneCClient) DoRequest(req *http.Request) ([]byte, *http.Response, error) {
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return nil, resp, fmt.Errorf("response error: code %d, body %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, resp, nil
}
