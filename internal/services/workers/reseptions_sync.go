package workers

import (
	"context"
	"log"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
)

type ReceptionRetryWorker struct {
	ReceptionRepo interfaces.ReceptionSmpRepository
	Hub           *websocket.Hub
	Interval      time.Duration
	Cancel        context.CancelFunc
	Logger        *log.Logger
}

func NewReceptionRetryWorker(
	repo interfaces.ReceptionSmpRepository,
	hub *websocket.Hub,
	interval time.Duration,
	logger *log.Logger,
) *ReceptionRetryWorker {
	return &ReceptionRetryWorker{
		ReceptionRepo: repo,
		Hub:           hub,
		Interval:      interval,
		Logger:        logger,
	}
}

// Start запускает воркер, который периодически пытается отправить недоставленные заявки
func (w *ReceptionRetryWorker) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	w.Cancel = cancel

	go func() {
		ticker := time.NewTicker(w.Interval)
		defer ticker.Stop()

		w.Logger.Printf("[ReceptionRetry] started, interval = %v", w.Interval)

		for {
			select {
			case <-ticker.C:
				w.Logger.Println("[ReceptionRetry] checking undelivered receptions...")
				if err := w.processUndeliveredReceptions(ctx); err != nil {
					w.Logger.Printf("[ReceptionRetry] processing failed: %v", err)
				} else {
					w.Logger.Println("[ReceptionRetry] undelivered receptions processed")
				}
			case <-ctx.Done():
				w.Logger.Println("[ReceptionRetry] stopped")
				return
			}
		}
	}()
}

// Stop завершает работу воркера
func (w *ReceptionRetryWorker) Stop() {
	if w.Cancel != nil {
		w.Cancel()
	}
}

// processUndeliveredReceptions — логика обработки недоставленных заявок
func (w *ReceptionRetryWorker) processUndeliveredReceptions(ctx context.Context) error {
	receptions, err := w.ReceptionRepo.GetUndeliveredReceptions(ctx)
	if err != nil {
		return err
	}

	for _, r := range receptions {
		var call models.Call
		// if err := json.Unmarshal(r.Data, &call); err != nil {
		// 	w.Logger.Printf("[ReceptionRetry] failed to unmarshal call %s: %v", r.CallID, err)
		// 	continue
		// }

		// Проверяем, онлайн ли доктор
		if w.Hub.IsUserConnected(call.Doctor.Id) {
			message := models.Message{
				// Type:   "reception_new",
				Header: "Новый вызов",
				Text:   "Поступил новый вызов",
				// Data:   call,
			}

			if w.Hub.SendToUserSafe(call.Doctor.Id, message) {
				// Успешно отправлено → обновляем статус
				if err := w.ReceptionRepo.UpdateStatus(ctx, r.CallID, "delivered"); err != nil {
					w.Logger.Printf("[ReceptionRetry] failed to update status to 'delivered' for %s: %v", r.CallID, err)
				}
			} else {
				// Не удалось отправить (буфер переполнен и т.п.)
				if err := w.ReceptionRepo.UpdateStatus(ctx, r.CallID, "failed"); err != nil {
					w.Logger.Printf("[ReceptionRetry] failed to update status to 'failed' for %s: %v", r.CallID, err)
				}
			}
		}
		// Если доктор не онлайн — оставляем статус как есть, попробуем позже
	}

	return nil
}
