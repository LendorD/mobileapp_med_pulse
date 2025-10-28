package usecases

import (
	"context"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/google/uuid"
)

type FlgUsecase struct {
	repo          interfaces.FlgRepository
	repoFile      interfaces.FileRepository
	repoTxManager interfaces.TxManager
	imageService  interfaces.ImageService
}

func NewFlgUseсase(repo interfaces.FlgRepository,
	fileRepo interfaces.FileRepository,
	txManager interfaces.TxManager,
	imageService interfaces.ImageService,
) interfaces.FlgUsecase {
	return &FlgUsecase{
		repo:          repo,
		repoFile:      fileRepo,
		imageService:  imageService,
		repoTxManager: txManager,
	}
}

func (u *FlgUsecase) CreateFlgWithPhoto(ctx context.Context, req *models.CreateFlgRequest) (*uint, *errors.AppError) {
	op := "usecase.Flg.CreateFlgWithPhoto"

	// Парсим дату (это не БД-операция, можно до транзакции)
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.NewAppError(400, "invalid date format, expected YYYY-MM-DD", err, true)
	}

	// Генерируем ключ для MinIO заранее (нужен для отката)
	minioKey := uuid.NewString()

	// === НАЧИНАЕМ ТРАНЗАКЦИЮ ===
	ctx, txErr := u.repoTxManager.Begin(ctx)
	if txErr != nil {
		return nil, errors.NewDBError(op, txErr)
	}

	// Флаг для контроля отката
	shouldRollback := true
	defer func() {
		if shouldRollback {
			// Сначала откатываем транзакцию
			_ = u.repoTxManager.Rollback(ctx)
			// Затем удаляем файл из MinIO (если он был загружен)
			_ = u.imageService.DeleteObject(ctx, minioKey)
		}
	}()

	// === 1. Загружаем файл в MinIO ===
	if err := u.imageService.UploadObject(ctx, minioKey, req.FileData); err != nil {
		return nil, errors.NewInternalError(op, "failed to upload image to storage", err)
	}

	// === 2. Создаём запись о файле в БД (в транзакции) ===
	file := &entities.File{
		MinioKey: minioKey,
		Filename: req.FileName,
	}
	if err := u.repoFile.CreateFile(ctx, file); err != nil {
		return nil, errors.NewDBError(op, err)
	}

	// === 3. Создаём FLG (в транзакции) ===
	flg := &entities.Flg{
		PatientID:    req.PatientID,
		Organization: req.Organization,
		Number:       req.Number,
		Result:       req.Result,
		Date:         parsedDate,
		FileID:       file.ID,
	}
	if err := u.repo.CreateFlg(ctx, flg); err != nil {
		return nil, errors.NewDBError(op, err)
	}

	// === УСПЕХ: отключаем откат ===
	shouldRollback = false

	// Коммитим транзакцию
	if err := u.repoTxManager.Commit(ctx); err != nil {
		// Если коммит не удался — файл уже в MinIO, но БД не сохранилась.
		// Удаляем файл, чтобы не было "висячих" объектов.
		_ = u.imageService.DeleteObject(ctx, minioKey)
		return nil, errors.NewDBError(op, err)
	}

	return &flg.ID, nil
}
