package auth

import (
	"context"
	"log"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"gorm.io/gorm"
)

func (r *AuthRepository) GetByLogin(ctx context.Context, login string) (*entities.Doctor, error) {
	op := "repo.Auth.GetByLogin"
	log.Printf("Searching for doctor with login: '%s'", login)

	var doctor entities.Doctor
	if err := r.db.WithContext(ctx).
		Where("phone = ?", login).
		First(&doctor).
		Error; err != nil {

		log.Printf("Error finding doctor: %v", err)
		return nil, errors.NewDBError(op, err)
	}

	log.Printf("Found doctor ID: %d", doctor.ID)
	return &doctor, nil
}

func (r *AuthRepository) SaveUsers(ctx context.Context, users []entities.AuthUser) error {
	tx := r.db.WithContext(ctx).Begin()
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
	err := r.db.WithContext(ctx).Where("login = ?", login).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
