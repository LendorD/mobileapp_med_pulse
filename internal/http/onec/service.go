package httpClient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func SendJSONRequest(httpMehod, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(httpMehod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки, если нужно
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Пример обработки ответа (смотрите, что необходимо для вашего случая)
	if resp.StatusCode != http.StatusOK {
		return respBody, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return respBody, nil
}
