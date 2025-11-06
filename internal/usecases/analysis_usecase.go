package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type AnalysisUsecase struct {
	analysisRepo      interfaces.AnalysisRepository
	analysisOrderRepo interfaces.AnalysisRepository
	txManager         interfaces.TxManager
}

func NewAnalysisUsecase(analysisRepo interfaces.AnalysisRepository, analysisOrderRepo interfaces.AnalysisRepository, txManager interfaces.TxManager) interfaces.AnalysisUsecase {
	return &AnalysisUsecase{
		analysisRepo:      analysisRepo,
		analysisOrderRepo: analysisOrderRepo,
		txManager:         txManager,
	}
}

func (u *AnalysisUsecase) UpdateAnalysisOrder(ctx context.Context, req *models.UpdateAnalysisOrderRequest) *errors.AppError {
	op := "usecase.AnalysisOrder.UpdateAnalysisOrder"

	// Начало транзакции
	ctx, err := u.txManager.Begin(ctx)
	if err != nil {
		return errors.NewDBError(op, err)
	}
	shouldRollback := true
	defer func() {
		if shouldRollback {
			_ = u.txManager.Rollback(ctx)
		}
	}()

	// 1. Проверка заказа
	order, appErr := u.analysisOrderRepo.GetAnalysisOrderByID(ctx, req.ID)
	if appErr != nil {
		return errors.NewDBError(op, appErr)
	}
	if order.PatientID != req.PatientID {
		return errors.NewForbiddenError(op, "order does not belong to patient")
	}

	// 2. Получение валидных анализов
	analysisIDs, appErr := u.analysisRepo.GetAllAnalysisIDs(ctx)
	if appErr != nil {
		return errors.NewDBError(op, appErr)
	}
	analysisMap := make(map[uint]entities.Analysis)
	for _, id := range analysisIDs {
		a, err := u.analysisRepo.GetAnalysisByID(ctx, id)
		if err != nil {
			continue
		}
		analysisMap[id] = *a
	}

	// 3. Валидация и подготовка item'ов
	now := time.Now()
	var itemsToUpsert []entities.AnalysisOrderItem
	for _, dto := range req.OrderItems {
		analysis, exists := analysisMap[dto.AnalysisID]
		if !exists {
			return errors.NewDBError(op, fmt.Errorf("analysis %d not found", dto.AnalysisID))
		}

		itemsToUpsert = append(itemsToUpsert, entities.AnalysisOrderItem{
			OrderID:           req.ID,
			AnalysisID:        dto.AnalysisID,
			IsCompleted:       dto.IsCompleted,
			CompletedAt:       dto.CompletedAt,
			PriceAtAssignment: analysis.Price,
			CreatedAt:         now,
			UpdatedAt:         now,
		})
	}

	// 4. Выполнение UPSERT
	if err := u.analysisOrderRepo.UpsertOrderItems(ctx, itemsToUpsert); err != nil {
		return errors.NewDBError(op, err)
	}

	// 5. Коммит
	shouldRollback = false
	if err := u.txManager.Commit(ctx); err != nil {
		return errors.NewDBError(op, err)
	}

	return nil
}

// GetAllAnalyses возвращает все доступные анализы
func (u *AnalysisUsecase) GetAllAnalyses(ctx context.Context) ([]models.AnalysisResponse, *errors.AppError) {

	analyses, err := u.analysisRepo.GetAllAnalyses(ctx)
	if err != nil {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to fetch analyses",
			err,
			false,
		)
	}

	result := make([]models.AnalysisResponse, len(analyses))
	for i, a := range analyses {
		result[i] = models.AnalysisResponse{
			ID:    a.ID,
			Code:  a.Code,
			Title: a.Title,
			Price: a.Price,
		}
	}

	return result, nil
}
