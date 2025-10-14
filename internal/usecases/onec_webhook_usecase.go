package usecases

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
	"github.com/gofrs/uuid"
)

type OneCWebhookUsecase struct {
	cacheRepo interfaces.OneCCacheRepository
	hub       *websocket.Hub
}

func NewOneCWebhookUsecase(
	cacheRepo interfaces.OneCCacheRepository,
	hub *websocket.Hub,
) interfaces.OneCWebhookUsecase {
	return &OneCWebhookUsecase{
		cacheRepo: cacheRepo,
		hub:       hub,
	}
}

// HandleReceptionsUpdate — обрабатывает обновление от 1С
func (u *OneCWebhookUsecase) HandleReceptionsUpdate(ctx context.Context, call models.Call) error {
	// 1. Сохраняем пациентов вызова в Redis по CallID (строка)
	err := u.cacheRepo.SaveReceptions(ctx, call.CallID, call.Patients)
	if err != nil {
		return err
	}

	// 2. Готовим сообщение для WebSocket-рассылки
	// Поскольку CallID — строка, а Message ожидает uint — у тебя два варианта:
	var typeID uint
	if callIDNum, err := strconv.Atoi(call.CallID); err == nil {
		typeID = uint(callIDNum)
	} else {
		// Если не число — используем 0 или хэш (но лучше переделать Message)
		typeID = 0
	}
	message := models.Message{
		Header:        "Новый вызов",
		Text:          fmt.Sprintf("Поступил вызов %s", call.CallID),
		TypeID:        typeID,      // uint (может быть 0, если CallID не число)
		Reference:     call.CallID, // ← строковый ID сохраняем здесь
		ReferenceID:   typeID,      // дублируем, если нужно
		GroupIDs:      nil,
		BroadcastUUID: uuid.Nil,
	}

	// 3. Отправляем broadcast-сообщение всем подписчикам
	u.hub.AddBroadcastMessage(message)

	return nil
}

// getInterestedUserIDs — приватный вспомогательный метод (не в интерфейсе!)
func (u *OneCWebhookUsecase) GetInterestedUserIDs(callID int) []uint {
	if callID == 123 {
		return []uint{1, 2, 3}
	}
	return []uint{1}
}
