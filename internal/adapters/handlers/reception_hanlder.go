package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // CreateReception godoc
// // @Summary Создать новый приём
// // @Description Создаёт запись о приёме пациента
// // @Tags Reception
// // @Accept json
// // @Produce json
// // @Param info body models.CreateReceptionRequest true "Данные приёма"
// // @Success 200 {object} entities.Reception "Созданный приём"
// // @Failure 400 {object} ResultError "Некорректный запрос"
// // @Failure 422 {object} ResultError "Ошибка валидации"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /reception [post]
// func (h *Handler) CreateReception(c *gin.Context) {
// 	var input models.CreateReceptionRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	if err := validate.Struct(input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	reception, eerr := h.usecase.Reception.Create(input)
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success reception create", apiresp.Object, reception)
// }

// // UpdateReceptionStatus godoc
// // @Summary Обновить статус приёма
// // @Description Изменяет статус приёма (scheduled, completed и т.д.)
// // @Tags Reception
// // @Accept json
// // @Produce json
// // @Param id path uint true "ID приёма"
// // @Param status body models.UpdateStatusRequest true "Новый статус"
// // @Success 200 {object} entities.Reception "Обновлённый приём"
// // @Failure 400 {object} ResultError "Некорректный запрос"
// // @Failure 404 {object} ResultError "Приём не найден"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /reception/{id}/status [put]
// func (h *Handler) UpdateReceptionStatus(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
// 		return
// 	}

// 	var input models.UpdateStatusRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	reception, eerr := h.usecase.Reception.UpdateStatus(uint(id), input.Status)
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success reception status update", apiresp.Object, reception)
// }

// // GetReceptionByID godoc
// // @Summary Получить приём по ID
// // @Description Возвращает информацию о приёме
// // @Tags Reception
// // @Accept json
// // @Produce json
// // @Param id path uint true "ID приёма"
// // @Success 200 {object} entities.Reception "Информация о приёме"
// // @Failure 400 {object} ResultError "Некорректный ID"
// // @Failure 404 {object} ResultError "Приём не найден"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /reception/{id} [get]
// func (h *Handler) GetReceptionByID(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
// 		return
// 	}

// 	reception, eerr := h.usecase.Reception.GetByID(uint(id))
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success reception get", apiresp.Object, reception)
// }

// GetReceptionsByDoctorAndDate godoc
// @Summary Получить записи врача на дату с пагинацией
// @Description Возвращает список записей для указанного врача на конкретную дату с сортировкой по статусу и пагинацией
// @Tags receptions
// @Produce json
// @Param doctor_id path int true "ID врача"
// @Param date query string true "Дата в формате YYYY-MM-DD"
// @Param page query int false "Номер страницы (начиная с 1)" default(1)
// @Success 200 {array} models.Reception
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /doctors/{doctor_id}/receptions [get]
func (h *Handler) GetReceptionsByDoctorAndDate(c *gin.Context) {
	// Получаем ID врача из URL
	doctorID, err := strconv.Atoi(c.Param("doctor_id"))
	if err != nil || doctorID < 1 { // Проверяем, что ID > 0
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID (must be positive integer)"})
		return
	}

	// Остальной код остаётся без изменений...
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date parameter is required"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be integer greater than 0"})
		return
	}

	receptions, err := h.usecase.GetReceptionsByDoctorAndDate(uint(doctorID), date, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}
