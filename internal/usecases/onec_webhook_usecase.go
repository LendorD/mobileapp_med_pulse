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
func (u *OneCWebhookUsecase) HandleReceptionsUpdate(DoctorID int, ctx context.Context, call models.Call) error {
	// err := u.repo.SaveReceptions(ctx, call.CallID, call.Patients)
	// if err != nil {
	// 	return err
	// }

	message := models.Message{
		Header: "Новый вызов",
		Text:   fmt.Sprintf("Поступил вызов %s", call.CallID),
	}

	u.hub.SendToUser(1, message)
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
