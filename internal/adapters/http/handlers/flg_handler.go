package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
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
	result := c.PostForm("result")
	date := c.PostForm("date")

	patientIDStr := c.PostForm("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 64)
	if err != nil || patientID == 0 {
		h.ErrorResponse(c, nil, http.StatusUnprocessableEntity, "invalid patient_id", true)
		return
	}

	// 2. Получение файла
	file, err := c.FormFile("file")
	if err != nil {
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
		PatientID: uint(patientID),
		Result:    result,
		Date:      date,
		FileName:  file.Filename,
		FileData:  data,
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
