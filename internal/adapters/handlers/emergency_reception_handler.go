package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
	"github.com/gin-gonic/gin"
)

////// TODO: нужно узнать по поводу закомменченных ручек: оставляем  или выкидываем?

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// CreateEmergencyReception godoc
// @Summary Создать экстренный приём
// @Description Создаёт запись об экстренном приёме
// @Tags EmergencyReception
// @Accept json
// @Produce json
// @Param info body models.CreateEmergencyRequest true "Данные экстренного приёма"
// @Success 200 {object} entities.EmergencyReception "Созданный экстренный приём"
// @Failure 400 {object} ResultError "Некорректный запрос"
// @Failure 422 {object} ResultError "Ошибка валидации"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /emergency [post]

// Примеры JSON
// Создание нового пациента на вызове
//
//	{
//	  "emergency_call_id": 123,
//	  "doctor_id": 1,
//	  "patient": {
//	    "full_name": "Иванов Иван Иванович",
//	    "birth_date": "1980-05-15",
//	    "is_male": true
//	  }
//	}
//
// Добавление существуещего пользователя
//
//	{
//	  "emergency_call_id": 124,
//	  "doctor_id": 2,
//	  "patient_id": 42
//	}
func (h *Handler) CreateEmergencyReception(c *gin.Context) {
	var input models.CreateEmergencyRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
		return
	}

	if err := validate.Struct(input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, errors.BadRequest, true)
		return
	}

	emergency, eerr := h.usecase.CreateReceptionSMP(&input)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success emergency reception create", Object, emergency)
}

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
//////

// GetEmergencyReceptionsByDoctorAndDate godoc
// @Summary Получить экстренные приёмы врача по дате
// @Description Возвращает список экстренных приёмов, назначенных врачу на указанную дату, с пагинацией
// @Tags EmergencyReception
// @Accept json
// @Produce json
// @Param doctor_id path uint true "ID врача"
// @Param date query string true "Дата в формате YYYY-MM-DD"
// @Param page query int false "Номер страницы" default(1)
// @Success 200 {array} entities.EmergencyReception "Список приёмов"
// @Failure 400 {object} ResultError "Некорректный запрос или параметры"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /emergency/{doctor_id}/receptions [get]
func (h *Handler) GetEmergencyReceptionsByDoctorAndDate(c *gin.Context) {
	// Получаем ID врача
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	// Получаем дату из query параметров
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

	// Получаем номер страницы
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be integer greater than 0"})
		return
	}

	// Вызываем usecase
	receptions, err := h.usecase.GetEmergencyReceptionsByDoctorAndDate(uint(doctorID), date, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}
