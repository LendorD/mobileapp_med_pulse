// handlers/onec.go

package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// OneCPatientListWebhook godoc
// @Summary Webhook from 1C: patient list update
// @Tags 1C
// @Accept json
// @Produce json
// @Param update body models.PatientListUpdate true "Patient list update from 1C"
// @Success 200
// @Router /webhook/onec/patients [post]
func (h *Handler) OneCPatientListWebhook(c *gin.Context) {
	var update models.PatientListUpdate
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
// @Summary Get cached patient list from 1C
// @Tags Patients
// @Produce json
// @Success 200 {object} models.PatientListResponse
// @Router /patients [get]
func (h *Handler) GetPatientList(c *gin.Context) {
	patients, err := h.usecase.GetPatientList(c.Request.Context())
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "Failed to get patient list from cache", true)
		return
	}

	response := models.PatientListResponse{
		Patient: patients,
		// Count:    len(patients),
	}

	h.ResultResponse(c, "success", Object, response)
}
