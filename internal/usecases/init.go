package usecases

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	_ "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
	_ "github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type UseCases struct {
	interfaces.MedCardUsecase
	interfaces.AuthUsecase
	interfaces.OneCWebhookUsecase
	interfaces.OneCPatientUsecase
}

func NewUsecases(
	r interfaces.Repository,
	s interfaces.Service,
	conf *config.Config,
	hub *websocket.Hub,
	onecClient interfaces.OneCClient,
) interfaces.Usecases {

	return &UseCases{
		NewMedCardUsecase(r, onecClient, r),
		NewAuthUsecase(r, conf.JWTSecret),
		NewOneCWebhookUsecase(r, hub),
		NewOneCPatientListUsecase(r, onecClient),
	}

}

// getFieldTypes возвращает карту, где ключ — это имя поля (по JSON-тегу),
// а значение — тип данных поля.
func getFieldTypes(model interface{}) (map[string]string, error) {
	result := make(map[string]string)

	// Получаем тип модели
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Разыменовываем указатель, если он есть
	}

	// Проверяем, что переданный объект является структурой
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", t.Kind())
	}

	// Итерируемся по полям структуры
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Получаем JSON-тег поля
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue // Пропускаем поля без JSON-тега
		}

		// Разбираем JSON-тег
		jsonName := strings.Split(jsonTag, ",")[0]

		// Получаем тип поля
		fieldType := field.Type

		// Если тип — указатель, получаем базовый тип
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		// Добавляем поле и его тип в карту
		result[jsonName] = fieldType.Name()
	}

	return result, nil
}
