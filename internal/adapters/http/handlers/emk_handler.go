package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllEmkByPatientID godoc
// @Summary Get all EMK records by patient ID
// @Tags EMK
// @Produce json
// @Param pat_id path string true "Patient ID"
// @Success 200 {array} entities.Emk
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security JWTAuth
// @Router /emk/{pat_id} [get]
func (h *Handler) GetAllEmkByPatientID(c *gin.Context) {
	patientID := c.Param("pat_id")
	if patientID == "" {
		h.ErrorResponse(c, http.ErrAbortHandler, http.StatusBadRequest, "patient_id is required", true)
		return
	}

	emk, err := h.usecase.GetAllEmkByPatientID(c.Request.Context(), patientID)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "medical card not found", true)
		return
	}
	h.ResultResponse(c, "success", Object, emk)
}
