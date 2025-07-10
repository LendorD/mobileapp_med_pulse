package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	authRepo  AuthRepository
	secretKey string
}

type AuthRepository interface {
	GetByLogin(ctx context.Context, login string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
}

func NewAuthUsecase(authRepo AuthRepository, secretKey string) *AuthUsecase {
	return &AuthUsecase{
		authRepo:  authRepo,
		secretKey: secretKey,
	}
}

func (uc *AuthUsecase) RegisterDoctor(ctx context.Context, req models.DoctorRegisterRequest) (*models.DoctorResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	existing, err := uc.authRepo.GetByLogin(ctx, req.Login)
	if err == nil && existing.ID != 0 {
		return nil, errors.New("doctor with this login already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		FullName:       req.FullName,
		Specialization: req.Specialization,
		Login:          req.Login,
		PasswordHash:   string(hashedPassword),
	}

	if err := uc.authRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &models.DoctorResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Specialization: user.Specialization,
	}, nil
}

func (uc *AuthUsecase) LoginDoctor(ctx context.Context, login, password string) (string, error) {
	user, err := uc.authRepo.GetByLogin(ctx, login)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(uc.secretKey))
}
