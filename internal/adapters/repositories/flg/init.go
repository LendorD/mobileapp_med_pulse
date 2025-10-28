package flg

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/base"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"gorm.io/gorm"
)

type FlgRepository struct {
	*base.BaseRepository
}

func NewFlgRepository(db *gorm.DB) interfaces.FlgRepository {
	return &FlgRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}
