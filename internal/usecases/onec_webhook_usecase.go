package usecases

import (
	"context"
	"fmt"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
)

type OneCWebhookUsecase struct {
	repo interfaces.ReceptionSmpRepository
	hub  *websocket.Hub
}

func NewOneCWebhookUsecase(
	repo interfaces.ReceptionSmpRepository,
	hub *websocket.Hub,
) interfaces.OneCWebhookUsecase {
	return &OneCWebhookUsecase{
		repo: repo,
		hub:  hub,
	}
}

// HandleReceptionsUpdate — обрабатывает обновление от 1С
func (u *OneCWebhookUsecase) HandleReceptionsUpdate(doctorID uint, ctx context.Context, call models.Call) error {
	// 1. Сохраняем заявку со статусом "received"
	err := u.repo.SaveReceptions(ctx, call.CallID, call)
	if err != nil {
		return err
	}

	// 2. Пытаемся отправить сразу
	if u.hub.IsUserConnected(doctorID) {
		message := models.Message{
			// Type:   "reception_new",
			Header: "Новый вызов",
			Text:   fmt.Sprintf("Поступил вызов %s", call.CallID),
			// Data:   call,
		}

		// Отправляем и меняем статус на "delivered"
		if u.hub.SendToUserSafe(doctorID, message) {
			// Успешно отправлено → обновляем статус
			return u.repo.UpdateStatus(ctx, call.CallID, "delivered")
		}
	}

	return nil
}

// getInterestedUserIDs — приватный вспомогательный метод (не в интерфейсе!)
func (u *OneCWebhookUsecase) GetInterestedUserIDs(callID int) []uint {
	if callID == 123 {
		return []uint{1, 2, 3}
	}
	return []uint{1}
}
