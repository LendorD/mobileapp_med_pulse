package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
)

type OneCWebhookUsecase struct {
	repo   interfaces.ReceptionSmpRepository
	hub    *websocket.Hub
	logger *logging.Logger
}

func NewOneCWebhookUsecase(
	repo interfaces.ReceptionSmpRepository,
	hub *websocket.Hub,
	logger *logging.Logger,
) interfaces.OneCWebhookUsecase {
	return &OneCWebhookUsecase{
		repo:   repo,
		hub:    hub,
		logger: logger,
	}
}

// HandleReceptionsUpdate — обрабатывает обновление от 1С
func (u *OneCWebhookUsecase) HandleReceptionsUpdate(doctorID uint, ctx context.Context, call models.Call) error {
	u.logger.Info("Начало обработки обновления вызова от 1С",
		"call_id", call.CallID,
		"doctor_id", doctorID,
	)

	// oneCReceptions, err := converter.CallToReception(call)
	// if err != nil {
	// 	return err
	// }

	// err = u.repo.SaveReceptions(ctx, call.CallID, oneCReceptions)
	// if err != nil {
	// 	return err
	// }

	callData, err := json.Marshal(call)
	if err != nil {
		u.logger.Error("Не удалось сериализовать вызов в JSON", "error", err)
		// Но всё равно отправим без данных
		callData = nil
	}

	// Пытаемся отправить сразу
	if u.hub.IsUserConnected(doctorID) {
		message := models.Message{
			Type:      "new_call",
			Header:    "Новый вызов",
			Text:      fmt.Sprintf("Поступил вызов %s", call.CallID),
			Data:      callData,
			CreatedAt: time.Now(),
		}
		u.logger.Debug("Сформировано уведомление для врача",
			"doctor_id", doctorID,
			"call_id", call.CallID,
			"message_header", message.Header,
		)

		// Отправляем и меняем статус на "delivered"
		if u.hub.SendToUserSafe(doctorID, message) {
			u.logger.Info("Уведомление успешно отправлено врачу через WebSocket",
				"doctor_id", doctorID,
				"call_id", call.CallID,
			)
			// Успешно отправлено → обновляем статус
			return u.repo.UpdateStatus(ctx, call.CallID, "delivered")
		} else {
			u.logger.Warn("Не удалось отправить уведомление врачу (SendToUserSafe вернул false)",
				"doctor_id", doctorID,
				"call_id", call.CallID,
			)
		}
	}
	u.logger.Info("Обработка вызова завершена",
		"call_id", call.CallID,
		"doctor_id", doctorID,
	)
	return nil
}

// getInterestedUserIDs — приватный вспомогательный метод (не в интерфейсе!)
func (u *OneCWebhookUsecase) GetInterestedUserIDs(callID int) []uint {
	if callID == 123 {
		return []uint{1, 2, 3}
	}
	return []uint{1}
}
