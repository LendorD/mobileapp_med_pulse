package usecases

import (
	"context"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type FileUsecase struct {
	repo         interfaces.FileRepository
	imageService interfaces.ImageService
}

func NewFileUsecase(repo interfaces.FileRepository, imageService interfaces.ImageService) interfaces.FileUsecase {
	return &FileUsecase{
		repo:         repo,
		imageService: imageService,
	}
}

func (u *FileUsecase) GetFileByID(ctx context.Context, fileID uint) ([]byte, string, string, *errors.AppError) {
	// 1. Юзкейс сам обращается к репозиторию
	fileMeta, err := u.repo.GetFileByID(ctx, fileID)
	if err != nil {
		return nil, "", "", &errors.AppError{
			Err:          err,
			Code:         http.StatusNotFound,
			Message:      "file not found",
			IsUserFacing: true,
		}
	}

	// 2. Передаёт данные в сервис
	data, contentType, err := u.imageService.GetFileByMinioKey(ctx, fileMeta.MinioKey, fileMeta.Filename)
	if err != nil {
		return nil, "", "", &errors.AppError{
			Err:          err,
			Code:         http.StatusInternalServerError,
			Message:      "failed to load file",
			IsUserFacing: false,
		}
	}

	return data, fileMeta.Filename, contentType, nil
}
