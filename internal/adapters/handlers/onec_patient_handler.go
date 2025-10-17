// handlers/onec.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// OneCPatientListWebhook godoc
// @Summary Webhook from 1C: patient list update
// @Tags 1C
// @Accept json
// @Produce json
// @Param update body entities.PatientListUpdate true "Patient list update from 1C"
// @Success 200
// @Router /webhook/onec/patients [post]
func (h *Handler) OneCPatientListWebhook(c *gin.Context) {
	var update entities.PatientListUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "Invalid patient list update payload", true)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.usecase.HandlePatientListUpdate(c.Request.Context(), update)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "Failed to handle patient list update", true)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	h.ResultResponse(c, "success", Empty, nil)
}

// GetPatientList godoc
// @Summary Get cached patient list from 1C with pagination
// @Tags Patients
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 20, max: 100)"
// @Success 200 {object} models.PatientListResponse
// @Router /patients [get]
func (h *Handler) GetPatientList(c *gin.Context) {
	// Получаем параметры пагинации
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 {
		limit = 20
	}
	if limit > 100 { // ограничим максимум
		limit = 100
	}

	offset := (page - 1) * limit

	// Получаем данные
	patients, err := h.usecase.GetPatientListPage(c.Request.Context(), offset, limit)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusInternalServerError, "Failed to get patient list from cache", true)
		return
	}

	response := models.PatientListResponse{
		Patient: patients,
	}

	h.ResultResponse(c, "success", Object, response)
}
