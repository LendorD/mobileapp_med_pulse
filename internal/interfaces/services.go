package interfaces

import (
	"context"
	"time"
)

type Service interface {
	ParamsParserService
	FilterBuilderService
	ImageService
}

// ParamsParserService Сервис преобразования типов
// Парсинг строковых параметров и приведение к единому типу
type ParamsParserService interface {
	ParseDateString(dateStr string) (time.Time, error)
	ParseTimeString(timeStr string) (time.Time, error)
	ParseUintString(uintStr string) (uint, error)
	ParseIntString(intStr string) (int, error)

	FormatDateToString(t time.Time) string
	FormatTimeToString(t time.Time) string
}

type FilterBuilderService interface {
	ParseFilterString(filterStr string, modelFields map[string]string) (string, []interface{}, error)
	ParseOrderString(orderStr string, modelFields map[string]string) (string, error)
}

type ImageService interface {
	UploadObject(ctx context.Context, key string, data []byte) error
	GetPresignedURL(ctx context.Context, key string) (string, error)
	DeleteObject(ctx context.Context, key string) error
	GetFileByMinioKey(ctx context.Context, minioKey, originalFilename string) ([]byte, string, error)
}
