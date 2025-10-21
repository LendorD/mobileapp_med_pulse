package workers

import (
	"context"
	"log"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
)

type PatientSyncWorker struct {
	Usecase  *usecases.OneCPatientUsecase
	Interval time.Duration
	Cancel   context.CancelFunc
}

func NewPatientSyncWorker(usecase *usecases.OneCPatientUsecase, interval time.Duration) *PatientSyncWorker {
	return &PatientSyncWorker{
		Usecase:  usecase,
		Interval: interval,
	}
}

// Start запускает воркер, который каждые N минут обновляет пациентов из 1С
func (w *PatientSyncWorker) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	w.Cancel = cancel

	go func() {
		ticker := time.NewTicker(w.Interval)
		defer ticker.Stop()

		log.Printf("[PatientSync] started, interval = %v", w.Interval)

		for {
			select {
			case <-ticker.C:
				log.Println("[PatientSync] updating patients from 1C...")
				if err := w.Usecase.UpdatePatientListFromOneC(ctx); err != nil {
					log.Printf("[PatientSync] update failed: %v", err)
				} else {
					log.Println("[PatientSync] patient list successfully updated")
				}
			case <-ctx.Done():
				log.Println("[PatientSync] stopped")
				return
			}
		}
	}()
}

// Stop завершает работу воркера
func (w *PatientSyncWorker) Stop() {
	if w.Cancel != nil {
		w.Cancel()
	}
}
