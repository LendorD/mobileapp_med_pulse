package usecases

import (
	"context"

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
func (u *OneCWebhookUsecase) HandleReceptionsUpdate(ctx context.Context, update models.OneCReceptionsUpdate) error {
	// 1. Сохраняем в Redis
	err := u.cacheRepo.SaveReceptions(ctx, update.CallID, update.Receptions)
	if err != nil {
		return err
	}

	// 2. Определяем, какие пользователи должны получить обновление
	// Например: все доктора, привязанные к этому вызову
	// Пока что — просто отправим всем (или по списку)
	// userIDs := u.GetInterestedUserIDs(update.CallID)

	// 3. Отправляем уведомление через WebSocket
	message := models.Message{
		Header:        "header",
		Text:          "text",
		TypeID:        uint(update.CallID),
		Reference:     "",
		ReferenceID:   uint(update.CallID),
		GroupIDs:      nil,
		BroadcastUUID: uuid.Nil,
	}

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
