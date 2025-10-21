package auth

import (
	"context"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

func (r *AuthRepository) SaveUsers(ctx context.Context, users []entities.AuthUser) error {
	db := r.db.GetDB(ctx)
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("DELETE FROM auth_users").Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(users) > 0 {
		if err := tx.CreateInBatches(users, 100).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *AuthRepository) GetUserByLogin(ctx context.Context, login string) (*entities.AuthUser, error) {
	var user entities.AuthUser
	db := r.db.GetDB(ctx)
	err := db.WithContext(ctx).Where("login = ?", login).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
