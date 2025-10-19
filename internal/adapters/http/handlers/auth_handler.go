package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// LoginDoctor аутентифицирует врача
// @Summary Вход в систему
// @Description Аутентифицирует врача по номеру телефона и паролю
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body models.DoctorLoginRequest true "Данные для входа"
// @Success 200 {object} models.DoctorAuthResponse "Успешное создание"
// @Failure 400 {object} IncorrectFormatError "Неверный формат запроса"
// @Failure 401 {object} IncorrectDataError "Неверные учётные данные"
// @Failure 500 {object} InternalServerError "Внутренняя ошибка сервера"
// @Router /auth [post]
func (h *Handler) LoginDoctor(c *gin.Context) {
	var req models.DoctorLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error decoding auth request", "error", err)
		h.ErrorResponse(c, err, http.StatusBadRequest, "Invalid request payload", true)
		return
	}

	h.logger.Info("Auth attempt", "phone", req.Phone)

	credentials, err := h.usecase.LoginDoctor(c.Request.Context(), req.Phone, req.Password)
	if err != nil {
		h.logger.Warn("Auth failed", "phone", req.Phone, "error", err)
		h.ErrorResponse(c, err, http.StatusBadRequest, "Invalid credentials", true)
		return
	}

	h.ResultResponse(c, "success", Object, credentials)
}

func (h *Handler) GetVersionProject(c *gin.Context) {
	version := "1.2.3"
	// version := h.usecase.GetVersion()
	c.JSON(http.StatusOK, gin.H{"version": version})
}
