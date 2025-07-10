package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
)

// AuthHandler обрабатывает запросы аутентификации
type AuthHandler struct {
	authUC *usecases.AuthUsecase
}

// NewAuthHandler создает новый экземпляр AuthHandler
func NewAuthHandler(authUC *usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

// RegisterDoctor регистрирует нового врача
// @Summary Регистрация врача
// @Description Регистрирует нового врача в системе
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body models.DoctorRegisterRequest true "Данные для регистрации"
// @Success 201 {object} models.DoctorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	var req models.DoctorRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	doctor, err := h.authUC.RegisterDoctor(r.Context(), req)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, doctor)
}

// LoginDoctor аутентифицирует врача
// @Summary Вход в систему
// @Description Аутентифицирует врача и возвращает JWT токен
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body models.DoctorLoginRequest true "Данные для входа"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) LoginDoctor(w http.ResponseWriter, r *http.Request) {
	var req models.DoctorLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := h.authUC.LoginDoctor(r.Context(), req.Login, req.Password)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
