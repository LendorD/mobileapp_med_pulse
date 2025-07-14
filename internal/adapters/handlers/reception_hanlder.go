package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// GetReceptionByID godoc
// @Summary Получить приём по ID
// @Description Возвращает информацию о приёме
// @Tags Reception
// @Accept json
// @Produce json
// @Param id path uint true "ID приёма"
// @Success 200 {object} entities.Reception "Информация о приёме"
// @Failure 400 {object} ResultError "Некорректный ID"
// @Failure 404 {object} ResultError "Приём не найден"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /reception/{id} [get]
func (h *Handler) GetReceptionsHospitalByPatientID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("pat_id"), 10, 64)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
		return
	}

	reception, eerr := h.usecase.GetReceptionsHospitalByPatientID(uint(id))
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success receptios get", Array, reception)
}

// GetPatientsByDoctorID godoc
// @Summary Получить приём по ID
// @Description Возвращает информацию о приёме
// @Tags Reception
// @Accept json
// @Produce json
// @Param id path uint true "ID приёма"
// @Success 200 {object} entities.Reception "Информация о приёме"
// @Failure 400 {object} ResultError "Некорректный ID"
// @Failure 404 {object} ResultError "Приём не найден"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /reception/{id} [get]
func (h *Handler) GetPatientsByDoctorID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("doc_id"), 10, 64)

	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
		return
	}
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		h.ErrorResponse(c, err, http.StatusBadRequest, "query parameter 'limit' must be a positive integer", false)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		h.ErrorResponse(c, err, http.StatusBadRequest, "query parameter 'offset' must be a non-negative integer", false)
		return
	}

	reception, eerr := h.usecase.GetPatientsByDoctorID(uint(id), limit, offset)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success receptios get", Array, reception)
}

// GetReceptionsByDoctorAndDate godoc
// @Summary Get receptions by doctor and date
// @Description Get paginated list of receptions for specific doctor and date
// @Tags receptions
// @Accept json
// @Produce json
// @Param doctor_id path int true "Doctor ID"
// @Param date query string false "Date in YYYY-MM-DD format"
// @Param page query int false "Page number" default(1)
// @Success 200 {array} models.ReceptionShortResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /receptions/doctor/{doctor_id} [get]
func (h *Handler) UpdateReceptionHospitalByReceptionID(c *gin.Context) {
	var input models.UpdateReceptionHospitalRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "Error create ReceptionHospitalRequest", true)
		return
	}

	if err := validate.Struct(input); err != nil {
		h.ErrorResponse(c, err, 422, "Error validate ReceptionHospitalRequest", true)
		return
	}

	doctor, eerr := h.usecase.UpdateReceptionHospital(&input)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}
	h.ResultResponse(c, "Success reception hospital update", Object, doctor)
}
