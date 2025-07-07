package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddPatientAllergy godoc
// @Summary Добавить аллергию пациенту
// @Description Добавляет аллергию в медицинскую карту пациента
// @Tags Allergy
// @Accept json
// @Produce json
// @Param info body models.AddAllergyRequest true "Данные аллергии"
// @Success 200 {object} entities.PatientsAllergy "Добавленная аллергия"
// @Failure 400 {object} ResultError "Некорректный запрос"
// @Failure 404 {object} ResultError "Пациент не найден"
// @Failure 422 {object} ResultError "Ошибка валидации"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /patient/allergy [post]
func (h *Handler) AddPatientAllergy(c *gin.Context) {
	var input models.AddAllergyRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	if err := validate.Struct(input); err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
		return
	}

	allergy, eerr := h.usecase.Allergy.AddToPatient(input.PatientID, input.AllergyID, input.Description)
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success allergy added", apiresp.Object, allergy)
}

// GetPatientAllergies godoc
// @Summary Получить аллергии пациента
// @Description Возвращает список аллергий пациента
// @Tags Allergy
// @Accept json
// @Produce json
// @Param patient_id path uint true "ID пациента"
// @Success 200 {array} entities.PatientsAllergy "Список аллергий"
// @Failure 400 {object} ResultError "Некорректный ID"
// @Failure 404 {object} ResultError "Пациент не найден"
// @Failure 500 {object} ResultError "Внутренняя ошибка"
// @Router /patient/{patient_id}/allergies [get]
func (h *Handler) GetPatientAllergies(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'patient_id' must be an integer", false)
		return
	}

	allergies, eerr := h.usecase.Allergy.GetByPatientID(uint(patientID))
	if eerr != nil {
		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
		return
	}

	h.ResultResponse(c, "Success patient allergies", apiresp.Array, allergies)
}
