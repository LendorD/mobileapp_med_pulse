package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // CreatePatient godoc
// // @Summary Создать нового пациента
// // @Description Создает нового пациента с персональными и контактными данными
// // @Tags Patient
// // @Accept json
// // @Produce json
// // @Param info body models.CreatePatientRequest true "Данные пациента"
// // @Success 200 {object} entities.Patient "Созданный пациент"
// // @Failure 400 {object} ResultError "Некорректный запрос"
// // @Failure 422 {object} ResultError "Ошибка валидации"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /patient [post]
// func (h *Handler) CreatePatient(c *gin.Context) {
// 	var input models.CreatePatientRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	if err := validate.Struct(input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	patient, eerr := h.usecase.Patient.Create(input)
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success patient create", apiresp.Object, patient)
// }

// // GetPatientByID godoc
// // @Summary Получить пациента по ID
// // @Description Возвращает полную информацию о пациенте
// // @Tags Patient
// // @Accept json
// // @Produce json
// // @Param id path uint true "ID пациента"
// // @Success 200 {object} entities.Patient "Информация о пациенте"
// // @Failure 400 {object} ResultError "Некорректный ID"
// // @Failure 404 {object} ResultError "Пациент не найден"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /patient/{id} [get]
// func (h *Handler) GetPatientByID(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
// 		return
// 	}

// 	patient, eerr := h.usecase.Patient.GetByID(uint(id))
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success patient get", apiresp.Object, patient)
// }

// // UpdatePatient godoc
// // @Summary Обновить данные пациента
// // @Description Обновляет информацию о пациенте
// // @Tags Patient
// // @Accept json
// // @Produce json
// // @Param info body models.UpdatePatientRequest true "Данные для обновления"
// // @Success 200 {object} entities.Patient "Обновленный пациент"
// // @Failure 400 {object} ResultError "Некорректный запрос"
// // @Failure 404 {object} ResultError "Пациент не найден"
// // @Failure 422 {object} ResultError "Ошибка валидации"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /patient [put]
// func (h *Handler) UpdatePatient(c *gin.Context) {
// 	var input models.UpdatePatientRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	if err := validate.Struct(input); err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, BadRequest, true)
// 		return
// 	}

// 	patient, eerr := h.usecase.Patient.Update(input)
// 	if eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success patient update", apiresp.Object, patient)
// }

// // DeletePatient godoc
// // @Summary Удалить пациента
// // @Description Удаляет пациента по ID
// // @Tags Patient
// // @Accept json
// // @Produce json
// // @Param id path uint true "ID пациента"
// // @Success 200 {object} ResultResponse "Успешное удаление"
// // @Failure 400 {object} ResultError "Некорректный ID"
// // @Failure 404 {object} ResultError "Пациент не найден"
// // @Failure 500 {object} ResultError "Внутренняя ошибка"
// // @Router /patient/{id} [delete]
// func (h *Handler) DeletePatient(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
// 		return
// 	}

// 	if eerr := h.usecase.Patient.Delete(uint(id)); eerr != nil {
// 		h.ErrorResponse(c, eerr.Err, eerr.Code, eerr.Message, eerr.IsUserFacing)
// 		return
// 	}

// 	h.ResultResponse(c, "Success patient delete", apiresp.Status, nil)
// }
