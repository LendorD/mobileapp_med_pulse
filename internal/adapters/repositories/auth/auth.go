package auth

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetByLogin(ctx context.Context, login string) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Where("login = ?", login).First(&user).Error
	return &user, err
}

func (r *AuthRepository) Create(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
