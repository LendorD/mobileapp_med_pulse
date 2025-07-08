package handlers

import (
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Object = "object"
	Array  = "array"
	Empty  = "empty"
)

type Response interface {
	ErrorResponse(c *gin.Context, err error, statusCode int, message string, showError bool)
	ResultResponse(c *gin.Context, message string, dataType string, data interface{})
}

func (h *Handler) ErrorResponse(c *gin.Context, err error, statusCode int, message string, showError bool) {
	errorMessage := message
	if showError && err != nil {
		errorMessage = message + ": " + err.Error()
	}

	c.JSON(statusCode, gin.H{
		"status": "error",
		"error": gin.H{
			"code":    statusCode,
			"message": errorMessage,
		},
	})
}

func (h *Handler) ResultResponse(c *gin.Context, message string, dataType string, data interface{}) {
	response := gin.H{
		"status":  "success",
		"message": message,
		"type":    dataType,
	}

	if data != nil {
		response["data"] = data
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) BadRequest(c *gin.Context, err error) {
	h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
}

func (h *Handler) InternalError(c *gin.Context, err error) {
	h.ErrorResponse(c, err, http.StatusInternalServerError, errors.InternalServerError, false)
}

func (h *Handler) NotFound(c *gin.Context, err error) {
	h.ErrorResponse(c, err, http.StatusNotFound, errors.NotFound, true)
}
