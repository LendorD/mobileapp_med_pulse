package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// GetMedCardByPatientID godoc
// @Summary Get medical card by patient ID
// @Tags MedicalCard
// @Produce json
// @Param pat_id path string true "Patient ID"
// @Success 200 {object} models.PatientCard
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /medcard/{pat_id} [get]
func (h *Handler) GetMedCardByPatientID(c *gin.Context) {
	patientID := c.Param("pat_id")
	if patientID == "" {
		h.ErrorResponse(c, http.ErrBodyNotAllowed, http.StatusBadRequest, "Failed to parse patientID", true)
		return
	}

	card, err := h.usecase.GetMedCardByPatientID(c.Request.Context(), patientID)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "Failed to fetch medical card", true)
		return
	}

	if card == nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "Medical card not found", true)
		return
	}
	h.ResultResponse(c, "success", Object, card)
}

// UpdateMedCard godoc
// @Summary Update medical card by patient ID
// @Tags MedicalCard
// @Accept json
// @Produce json
// @Param pat_id path string true "Patient ID"
// @Param update body models.UpdateMedicalCardRequest true "Medical card update data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /medcard/{pat_id} [put]
func (h *Handler) UpdateMedCard(c *gin.Context) {
	patientID := c.Param("pat_id")
	if patientID == "" {
		h.ErrorResponse(c, http.ErrBodyNotAllowed, http.StatusBadRequest, "Failed to parse patientID", true)
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient_id is required"})
		return
	}

	var req models.UpdateMedicalCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ErrorResponse(c, http.ErrBodyNotAllowed, http.StatusBadRequest, "Invalid request body", true)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Убеждаемся, что ID из URL совпадает с ID в теле (опционально, но безопасно)
	if req.PatientID != patientID {
		h.ErrorResponse(c, http.ErrBodyNotAllowed, http.StatusBadRequest, "Patient_id in URL and body must match", true)
		return
	}

	err := h.usecase.UpdateMedicalCard(c.Request.Context(), &req)
	if err != nil {
		h.ErrorResponse(c, http.ErrBodyNotAllowed, http.StatusBadRequest, "Failed to update medical card", true)
		return
	}

	h.ResultResponse(c, "success", Empty, nil)
}
