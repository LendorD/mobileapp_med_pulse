package onec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
)

// Client — HTTP-клиент для взаимодействия с 1С
type Client struct {
	baseURL    string
	httpClient *http.Client
	username   string
	password   string
}

// NewClient создаёт новый клиент для 1С
func NewClient(cfg config.OneCConfig) *Client {
	// Создаём HTTP-клиент с таймаутом
	client := &http.Client{
		Timeout: cfg.Timeout,
	}

	return &Client{
		baseURL:    cfg.BaseURL,
		httpClient: client,
		username:   cfg.Username,
		password:   cfg.Password,
	}
}

// buildURL собирает полный URL: baseURL + path
func (c *Client) buildURL(path string) string {
	// Убираем завершающий слеш у baseURL и добавляем путь
	base := c.baseURL
	if base[len(base)-1] == '/' {
		base = base[:len(base)-1]
	}
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return base + path
}

// do выполняет HTTP-запрос с базовой аутентификацией (если заданы логин/пароль)
func (c *Client) do(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	url := c.buildURL(path)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Базовая аутентификация, если указаны учётные данные
	if c.username != "" || c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	// Проверяем статус-код
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("1C returned status %d: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

// Get выполняет GET-запрос к 1С и возвращает тело ответа
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	resp, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

// Post отправляет JSON в 1С и возвращает ответ
func (c *Client) Post(ctx context.Context, path string, data interface{}) ([]byte, error) {
	var bodyBytes []byte
	var err error

	if data != nil {
		bodyBytes, err = json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request data: %w", err)
		}
	}

	resp, err := c.do(ctx, http.MethodPost, path, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseBody, nil
}
