package models

// CreateFlgRequest — данные для создания FLG (уже без multipart!)
type CreateFlgRequest struct {
	PatientID    uint   `json:"patient_id"`
	Organization string `json:"organization"`
	Number       string `json:"number"`
	Result       string `json:"result"`
	Date         string `json:"date"` // "2025-10-14"
	FileName     string `json:"file_name"`

	// Данные изображения
	FileData []byte `json:"-"` // не сериализуется в JSON
}

// FlgResponse — ответ
type FlgResponse struct {
	ID           uint   `json:"id"`
	Organization string `json:"organization"`
	Number       string `json:"number"`
	Result       string `json:"result"`
	Date         string `json:"date"`
	PhotoURL     string `json:"photo_url"` // minio_id каждый раз проксировать и брать из минио а не хранить ссылку! сделать отдельную таблицу для файлов
}
