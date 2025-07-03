package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

// LoginRequest represents login credentials
// @Description Login credentials structure
type LoginRequest struct {
	Login    string `json:"login" binding:"required" example:"doctor_ivanov"`
	Password string `json:"password" binding:"required" example:"securepassword123"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new doctor
// @Description Register a new doctor with the input payload
// @Tags auth
// @Accept  json
// @Produce  json
// @Param doctor body models.Doctor true "Doctor information"
// @Success 201 {object} models.Doctor
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Register(&doctor); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "doctor with this login already exists" {
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	doctor.PasswordHash = ""
	c.JSON(http.StatusCreated, doctor)
}

// Login godoc
// @Summary Login a doctor
// @Description Login with login and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body LoginRequest true "Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var credentials LoginRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(credentials.Login, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
