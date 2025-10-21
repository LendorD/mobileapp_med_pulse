package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/gin-gonic/gin"
)

// SaveSignature uploads a signature image for a given reception ID.
// @Summary Upload patient signature
// @Description Accepts a multipart form with a 'signature' file (e.g., PNG, JPG) and saves it.
// @Tags Emergency
// @Accept multipart/form-data
// @Produce json
// @Param recep_id path string true "Reception ID"
// @Param signature formData file true "Signature image file"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /signature/{recep_id} [post]
func (h *Handler) SaveSignature(c *gin.Context) {
	// patientID, err := h.service.ParseUintString(c.Param("recep_id"))
	// if err != nil {
	// 	h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
	// 	return
	// }

	// читаем файл из multipart
	file, err := c.FormFile("signature")
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		h.ErrorResponse(c, err, http.StatusInternalServerError, errors.InternalServerError, false)
		return
	}
	defer openedFile.Close()

	// читаем все байты
	// signatureBytes, err := io.ReadAll(openedFile)
	// if err != nil {
	// 	h.ErrorResponse(c, err, http.StatusInternalServerError, errors.InternalServerError, false)
	// 	return
	// }

	// сохраняем через usecase
	// if appErr := h.usecase.SavePatientSignature(patientID, signatureBytes); appErr != nil {
	// 	h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
	// 	return
	// }

	h.ResultResponse(c, "Signature saved", Object, nil)
}

// GetSignature retrieves the patient's signature by reception ID.
// @Summary Get patient signature
// @Description Returns the base64-encoded signature image for a given reception ID.
// @Tags Emergency
// @Produce json
// @Param recep_id path string true "Reception ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /signature/{recep_id} [get]
func (h *Handler) GetSignature(c *gin.Context) {
	// patientID, err := h.service.ParseUintString(c.Param("recep_id"))
	// if err != nil {
	// 	h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
	// 	return
	// }

	// // signature, appErr := h.usecase.GetPatientSignature(patientID)
	// if appErr != nil {
	// 	h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
	// 	return
	// }

	// h.ResultResponse(c, "Signature fetched", Object, gin.H{"signatureBase64": signature})
}
