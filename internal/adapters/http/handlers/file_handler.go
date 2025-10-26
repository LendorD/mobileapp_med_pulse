package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFlgWithPhoto godoc
// @Summary Создать флюорографию с фото
// @Description Загружает фото и создаёт запись
// @Tags Flg
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param patient_id formData uint true "ID пациента"
// @Param organization formData string true "Организация"
// @Param number formData string true "Номер"
// @Param result formData string true "Результат"
// @Param date formData string true "Дата (YYYY-MM-DD)"
// @Param file formData file true "Фото (JPEG/PNG, до 10 МБ)"
// @Success 201 {object} ResultResponse{models.FlgResponse} "Созданная флюрография"
// @Failure 400 {object} ResultError "Некорректный запрос (например, неверный формат даты)"
// @Failure 401 {object} ResultError "Неавторизован (отсутствует или невалиден токен)"
// @Failure 404 {object} ResultError "Пациент или организация не найдены"
// @Failure 422 {object} ResultError "Семантическая ошибка (например, patient_id=0)"
// @Failure 500 {object} ResultError "Внутренняя ошибка сервера"
// @Router /flgs/ [post]
func (h *Handler) CreateFlgWithPhoto(c *gin.Context) {
	// 1. Парсинг скалярных полей
	organization := c.PostForm("organization")
	number := c.PostForm("number")
	result := c.PostForm("result")
	date := c.PostForm("date")

	patientIDStr := c.PostForm("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 64)
	if err != nil || patientID == 0 {
		h.ErrorResponse(c, nil, http.StatusUnprocessableEntity, "invalid patient_id", true)
		return
	}

	// 2. Получение файла (гарантировано существует благодаря мидлвари)
	file, err := c.FormFile("file")
	if err != nil {
		// На самом деле, сюда не должно дойти — мидлварь уже проверил
		h.ErrorResponse(c, err, http.StatusBadRequest, "file is required", true)
		return
	}

	// 3. Чтение данных
	src, err := file.Open()
	if err != nil {
		h.ErrorResponse(c, err, http.StatusInternalServerError, "failed to open file", false)
		return
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusInternalServerError, "failed to read file", false)
		return
	}

	// 4. Формируем запрос для юзкейса
	req := &models.CreateFlgRequest{
		PatientID:    uint(patientID),
		Organization: organization,
		Number:       number,
		Result:       result,
		Date:         date,
		FileName:     file.Filename,
		FileData:     data,
		// ContentType можно не передавать — выведем из Filename в юзкейсе или репозитории
	}

	// 5. Вызов юзкейса
	resp, appErr := h.usecase.CreateFlgWithPhoto(c.Request.Context(), req)
	if appErr != nil {
		h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Flg created sucessfully", Object, resp)
}

// UpdateAnalysisOrder godoc
// @Summary Обновить направление на анализы
// @Description Обновляет список анализов в направлении
// @Tags AnalysisOrder
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param info body models.UpdateAnalysisOrderRequest true "Данные направления"
// @Success 204 "Направление успешно обновлено"
// @Failure 400 {object} ResultError "Неверный формат запроса"
// @Failure 422 {object} ResultError "Ошибка валидации данных"
// @Failure 404 {object} ResultError "Направление не найдено"
// @Failure 500 {object} ResultError "Внутренняя ошибка сервера"
// @Router /analysis/update [post]
func (h *Handler) UpdateAnalysisOrder(c *gin.Context) {
	var req models.UpdateAnalysisOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "invalid request body", true)
		return
	}

	appErr := h.usecase.UpdateAnalysisOrder(c.Request.Context(), &req)
	if appErr != nil {
		h.ErrorResponse(c, appErr.Err, appErr.Code, appErr.Message, appErr.IsUserFacing)
		return
	}

	c.Status(http.StatusNoContent)
}

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
