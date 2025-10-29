package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetFileByID godoc
// @Summary Получить файл по ID
// @Description Возвращает бинарные данные файла по его идентификатору. Устанавливает правильные заголовки Content-Type, Content-Length и Content-Disposition.
// @Tags Files
// @Produce octet-stream
// @Security BearerAuth
// @Param id path uint true "ID файла"
// @Success 200 {string} binary "Файл в виде бинарных данных"
// @Failure 400 {object} ResultError "Некорректный ID файла (не число)"
// @Failure 401 {object} ResultError "Неавторизован (отсутствует или невалиден токен)"
// @Failure 404 {object} ResultError "Файл не найден"
// @Failure 500 {object} ResultError "Внутренняя ошибка сервера"
// @Router /files/{id} [get]
func (h *Handler) GetFileByID(c *gin.Context) {
	fileIDStr := c.Param("id")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		h.ErrorResponse(c, nil, http.StatusBadRequest, "invalid file ID", true)
		return
	}

	fileData, filename, contentType, appErr := h.usecase.GetFileByID(c.Request.Context(), uint(fileID))
	if appErr != nil {
		h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
		return
	}

	// ✅ Используем contentType, который пришёл из юзкейса
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", strconv.Itoa(len(fileData)))
	c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")

	c.Data(http.StatusOK, contentType, fileData)
}
