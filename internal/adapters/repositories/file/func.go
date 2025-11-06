package file

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

// CreateFile создаёт запись о файле в БД
func (r *FileRepository) CreateFile(ctx context.Context, file *entities.File) error {
	op := "repo.File.CreateFile"
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Create(file).Error; err != nil {
		return errors.NewDBError(op, err)
	}
	return nil
}

// GetFileByID возвращает файл по ID
func (r *FileRepository) GetFileByID(ctx context.Context, id uint) (*entities.File, error) {
	op := "repo.File.GetFileByID"
	var file entities.File
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Where("id = ?", id).First(&file).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFoundError(op)
		}
		return nil, errors.NewDBError(op, err)
	}
	return &file, nil
}

// DeleteFile удаляет файл по ID (жёсткое удаление)
func (r *FileRepository) DeleteFile(ctx context.Context, id uint) error {
	op := "repo.File.DeleteFile"
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Where("id = ?", id).Delete(&entities.File{}).Error; err != nil {
		return errors.NewDBError(op, err)
	}
	return nil
}
