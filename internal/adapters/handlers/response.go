package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/gin-gonic/gin"
)

// Константы типов ответа
const (
	Object = "object" // Используется когда ответ содержит один объект
	Array  = "array"  // Используется когда ответ содержит массив объектов
	Empty  = "empty"  // Используется когда ответ не содержит данных
)

// Response - интерфейс для стандартных ответов API
type Response interface {
	ErrorResponse(c *gin.Context, err error, statusCode int, message string, showError bool)
	ResultResponse(c *gin.Context, message string, dataType string, data interface{})
}

// ErrorResponse - возвращает стандартизированный ответ с ошибкой
// @Summary Ответ с ошибкой
// @Description Возвращает JSON с деталями ошибки
// @Param c body context.Context true "Контекст Gin"
// @Param err error true "Объект ошибки"
// @Param statusCode int true "HTTP статус код"
// @Param message string true "Сообщение об ошибке"
// @Param showError bool true "Показывать ли детали ошибки"
// @Failure 400 {object} map[string]interface{} "Неверный запрос"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
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

// ResultResponse - возвращает стандартизированный успешный ответ
// @Summary Успешный ответ
// @Description Возвращает JSON с данными
// @Param c body context.Context true "Контекст Gin"
// @Param message string true "Сообщение для клиента"
// @Param dataType string true "Тип данных (object/array/empty)"
// @Param data interface{} false "Данные для возврата"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
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

// BadRequest - возвращает ошибку 400
// @Summary Неверный запрос
// @Description Возвращает ошибку 400 Bad Request
func (h *Handler) BadRequest(c *gin.Context, err error) {
	h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
}

// InternalError - возвращает ошибку 500
// @Summary Ошибка сервера
// @Description Возвращает ошибку 500 Internal Server Error
func (h *Handler) InternalError(c *gin.Context, err error) {
	h.ErrorResponse(c, err, http.StatusInternalServerError, errors.InternalServerError, false)
}

// NotFound - возвращает ошибку 404
// @Summary Не найдено
// @Description Возвращает ошибку 404 Not Found
func (h *Handler) NotFound(c *gin.Context, err error) {
	h.ErrorResponse(c, err, http.StatusNotFound, errors.NotFound, true)
}

// RespondWithError - вспомогательная функция для возврата ошибки (для чистого http)
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON - вспомогательная функция для возврата JSON (для чистого http)
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
