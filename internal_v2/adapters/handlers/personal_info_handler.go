package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPersonalInfoByPatientID godoc
// @Summary Получить персональные данные пациента
// @Description Возвращает персональные данные пациента по ID пациента
// @Tags PersonalInfo
// @Accept json
// @Produce json
// @Param patient_id path uint true "ID пациента"
// @Success 200 {object} entities.PersonalInfo "Персональные данные"
// @Failure 400 {object} ResultError "Некорректный ID пациента"
// @Failure 404 {object} ResultError "Данные не найдены"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /patient/{patient_id}/personal-info [get]
func (h *Handler) GetPersonalInfoByPatientID(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'patient_id' must be an integer", false)
		return
	}

	info, eerr := h.usecase.PersonalInfo.GetByPatientID(uint(patientID))
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success personal info", apiresp.Object, info)
}

// UpdatePersonalInfo godoc
// @Summary Обновить персональные данные
// @Description Обновляет персональные данные пациента
// @Tags PersonalInfo
// @Accept json
// @Produce json
// @Param patient_id path uint true "ID пациента"
// @Param info body models.UpdatePersonalInfoRequest true "Данные для обновления"
// @Success 200 {object} entities.PersonalInfo "Обновленные данные"
// @Failure 400 {object} ResultError "Некорректный запрос"
// @Failure 404 {object} ResultError "Пациент не найден"
// @Failure 422 {object} ResultError "Ошибка валидации"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /patient/{patient_id}/personal-info [put]
func (h *Handler) UpdatePersonalInfo(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'patient_id' must be an integer", false)
		return
	}

	var input models.UpdatePersonalInfoRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	input.PatientID = uint(patientID)

	if err := validate.Struct(input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	info, eerr := h.usecase.PersonalInfo.Update(input)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success personal info update", apiresp.Object, info)
}
