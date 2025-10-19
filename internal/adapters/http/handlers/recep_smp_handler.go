package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/gin-gonic/gin"
)

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

func (h *Handler) GetPdf(c *gin.Context) {
	fmt.Println("Зашли для получения pdf")
	dir, err := os.Getwd() // текущая рабочая директория, обычно это cmd/app
	fmt.Println("директория:", dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get working directory"})
		return
	}

	pdfPath := filepath.Join(dir, "assets", "mobileapp.pdf")

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		h.ErrorResponse(c, err, http.StatusBadRequest, "PDF not found", true)
		return
	}

	c.File(pdfPath) // Gin сам установит нужные заголовки
}
