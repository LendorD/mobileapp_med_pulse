package handlers

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
