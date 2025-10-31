package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// UpdateAnalysisOrder godoc
// @Summary Обновить направление на анализы
// @Description Обновляет список анализов в направлении
// @Tags Files
// @Accept json
// @Produce json
// @Param info body models.UpdateAnalysisOrderRequest true "Данные направления"
// @Success 204 "Направление успешно обновлено"
// @Failure 400 {object} ResultError "Неверный формат запроса"
// @Failure 422 {object} ResultError "Ошибка валидации данных"
// @Failure 404 {object} ResultError "Направление не найдено"
// @Failure 500 {object} ResultError "Внутренняя ошибка сервера"
// @Router /analysis/update [post]
// @Security JWTAuth
func (h *Handler) UpdateAnalysisOrder(c *gin.Context) {
	var req models.UpdateAnalysisOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "invalid request body", true)
		return
	}

	appErr := h.usecase.UpdateAnalysisOrder(c.Request.Context(), &req)
	if appErr != nil {
		h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllAnalyses godoc
// @Summary Получить все анализы
// @Description Возвращает полный список доступных анализов
// @Tags Files
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.AnalysisResponse "Список анализов"
// @Failure 500 {object} ResultError "Внутренняя ошибка сервера"
// @Router /analysis [get]
// @Security JWTAuth
func (h *Handler) GetAllAnalyses(c *gin.Context) {
	analyses, appErr := h.usecase.GetAllAnalyses(c.Request.Context())
	if appErr != nil {
		h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Analysis retrieved successfully", Array, analyses)
}
