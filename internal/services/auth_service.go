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
	repo         *repository.AuthRepository
	jwtSecret    string
	accessExpiry time.Duration
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func NewAuthService(repo *repository.AuthRepository, secret string, accessExpiry time.Duration) *AuthService {
	return &AuthService{
		repo:         repo,
		jwtSecret:    secret,
		accessExpiry: accessExpiry,
	}
}

func (s *AuthService) Register(doctor *models.Doctor) error {
	// Проверка уникальности логина
	if existing, _ := s.repo.FindDoctorByLogin(doctor.Login); existing != nil {
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

func (s *AuthService) Login(login, password string) (*TokenPair, error) {
	doctor, err := s.repo.FindDoctorByLogin(login)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(doctor.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Генерация токенов
	accessToken, err := s.generateAccessToken(doctor.ID)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) generateAccessToken(doctorID uint) (string, error) {
	claims := jwt.MapClaims{
		"doctor_id": doctorID,
		"exp":       time.Now().Add(s.accessExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	doctorID := uint(claims["doctor_id"].(float64))
	return doctorID, nil
}
