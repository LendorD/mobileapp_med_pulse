package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/models"
	"github.com/gin-gonic/gin"
)

// CreateDoctor godoc
// @Summary Создать нового врача
// @Description Создает нового врача с указанными данными
// @Tags Doctor
// @Accept json
// @Produce json
// @Param info body models.CreateDoctorRequest true "Данные врача"
// @Success 200 {object} entities.Doctor "Созданный врач"
// @Failure 400 {object} ResultError "Некорректный запрос"
// @Failure 422 {object} ResultError "Ошибка валидации"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /doctor [post]
func (h *Handler) CreateDoctor(c *gin.Context) {
	var input models.CreateDoctorRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	if err := validate.Struct(input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	doctor, eerr := h.usecase.Doctor.Create(input)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success doctor create", apiresp.Object, doctor)
}

// GetDoctorByID godoc
// @Summary Получить врача по ID
// @Description Возвращает информацию о враче по ID
// @Tags Doctor
// @Accept json
// @Produce json
// @Param id path uint true "ID врача"
// @Success 200 {object} entities.Doctor "Информация о враче"
// @Failure 400 {object} ResultError "Некорректный ID"
// @Failure 404 {object} ResultError "Врач не найден"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /doctor/{id} [get]
func (h *Handler) GetDoctorByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
		return
	}

	doctor, eerr := h.usecase.Doctor.GetByID(uint(id))
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success doctor get", apiresp.Object, doctor)
}

// UpdateDoctor godoc
// @Summary Обновить данные врача
// @Description Обновляет информацию о враче
// @Tags Doctor
// @Accept json
// @Produce json
// @Param info body models.UpdateDoctorRequest true "Данные для обновления"
// @Success 200 {object} entities.Doctor "Обновленный врач"
// @Failure 400 {object} ResultError "Некорректный запрос"
// @Failure 404 {object} ResultError "Врач не найден"
// @Failure 422 {object} ResultError "Ошибка валидации"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /doctor [put]
func (h *Handler) UpdateDoctor(c *gin.Context) {
	var input models.UpdateDoctorRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	if err := validate.Struct(input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	doctor, eerr := h.usecase.Doctor.Update(input)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success doctor update", apiresp.Object, doctor)
}

// DeleteDoctor godoc
// @Summary Удалить врача
// @Description Удаляет врача по ID
// @Tags Doctor
// @Accept json
// @Produce json
// @Param id path uint true "ID врача"
// @Success 200 {object} ResultResponse "Успешное удаление"
// @Failure 400 {object} ResultError "Некорректный ID"
// @Failure 404 {object} ResultError "Врач не найден"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /doctor/{id} [delete]
func (h *Handler) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
		return
	}

	if eerr := h.usecase.Doctor.Delete(uint(id)); eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success doctor delete", apiresp.Status, nil)
}

type DoctorFilterResponse = models.FilterResponse[[]entities.Doctor]

// GetFilteredDoctors godoc
// @Summary Получить отфильтрованный список врачей
// @Description Возвращает список врачей с фильтрацией и пагинацией
// @Tags Doctor
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param count query int false "Количество записей" default(10)
// @Param order query string false "Сортировка"
// @Param filter query string false "Фильтрация"
// @Success 200 {object} DoctorFilterResponse "Список врачей"
// @Failure 400 {object} ResultError "Некорректные параметры"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /doctor [get]
func (h *Handler) GetFilteredDoctors(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'page' must be an integer", false)
		return
	}

	count, err := strconv.Atoi(c.DefaultQuery("count", "0"))
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'count' must be an integer", false)
		return
	}

	order := c.Query("order")
	filter := c.Query("filter")

	doctors, eerr := h.usecase.Doctor.GetFiltered(page, count, order, filter)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success get doctors", apiresp.Array, doctors)
}
