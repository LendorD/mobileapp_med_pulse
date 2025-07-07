package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // CreateEmergencyReception godoc
// // @Summary Создать экстренный приём
// // @Description Создаёт запись об экстренном приёме
// // @Tags EmergencyReception
// // @Accept json
// // @Produce json
// // @Param info body models.CreateEmergencyRequest true "Данные экстренного приёма"
// // @Success 200 {object} entities.EmergencyReception "Созданный экстренный приём"
// // @Failure 400 {object} ResultError "Некорректный запрос"
// // @Failure 422 {object} ResultError "Ошибка валидации"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /emergency [post]
// func (h *Handler) CreateEmergencyReception(c *gin.Context) {
// 	var input models.CreateEmergencyRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	if err := validate.Struct(input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	emergency, eerr := h.usecase.Emergency.Create(input)
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success emergency reception create", apiresp.Object, emergency)
// }

// // AssignDoctorToEmergency godoc
// // @Summary Назначить врача на экстренный приём
// // @Description Назначает врача на экстренный вызов
// // @Tags EmergencyReception
// // @Accept json
// // @Produce json
// // @Param id path uint true "ID экстренного приёма"
// // @Param doctor_id body models.AssignDoctorRequest true "ID врача"
// // @Success 200 {object} entities.EmergencyReception "Обновлённый экстренный приём"
// // @Failure 400 {object} ResultError "Некорректный запрос"
// // @Failure 404 {object} ResultError "Ресурс не найден"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /emergency/{id}/assign [put]
// func (h *Handler) AssignDoctorToEmergency(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
// 		return
// 	}

// 	var input models.AssignDoctorRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	emergency, eerr := h.usecase.Emergency.AssignDoctor(uint(id), input.DoctorID)
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success doctor assigned to emergency", apiresp.Object, emergency)
// }
