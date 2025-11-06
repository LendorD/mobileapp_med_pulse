package file

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type FileRepository struct {
	*base.BaseRepository
}

func NewFileRepository(db *gorm.DB) interfaces.FileRepository {
	return &FileRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}
