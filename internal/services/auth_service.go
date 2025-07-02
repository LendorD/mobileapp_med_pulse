package services

import (
	"errors"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo      *repository.AuthRepository
	jwtSecret string
}

func NewAuthService(repo *repository.AuthRepository, secret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: secret}
}

func (s *AuthService) Register(doctor *models.Doctor) error {
	// Проверка уникальности логина
	existing, _ := s.repo.FindDoctorByLogin(doctor.Login)
	if existing != nil {
		return errors.New("doctor with this login already exists")
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(doctor.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	doctor.PasswordHash = string(hashedPassword)

	return s.repo.CreateDoctor(doctor)
}

func (s *AuthService) Login(login, password string) (string, error) {
	doctor, err := s.repo.FindDoctorByLogin(login)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(doctor.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Генерация JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"doctor_id": doctor.ID,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return uint(claims["doctor_id"].(float64)), nil
}
