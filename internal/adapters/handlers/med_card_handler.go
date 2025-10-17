package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "patient_id is required"})
		return
	}

	card, err := h.usecase.GetMedCardByPatientID(c.Request.Context(), patientID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "medical card not found"})
		return
	}

	c.JSON(http.StatusOK, card) // ← entities.OneCMedicalCard с тегами json
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "patient_id is required"})
		return
	}

	var card entities.OneCMedicalCard
	if err := c.ShouldBindJSON(&card); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Убеждаемся, что ID из URL совпадает с ID в теле
	if card.PatientID != patientID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "patient_id in URL and body must match"})
		return
	}

	if err := h.usecase.UpdateMedicalCard(c.Request.Context(), &card); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update medical card"})
		return
	}

	c.Status(http.StatusOK)
}
