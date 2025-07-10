package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
)

type AuthHandler struct {
	authUC *usecases.AuthUsecase
}

func NewAuthHandler(authUC *usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

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
