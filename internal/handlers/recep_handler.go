package handlers


import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/gin-gonic/gin"
)

type ReceptionHandler struct {
	receptionService services.ReceptionService
}

func NewReceptionHandler(receptionService services.ReceptionService) *ReceptionHandler {
	return &ReceptionHandler{receptionService: receptionService}
}

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
func (h *ReceptionHandler) GetReceptionsByDoctorAndDate(c *gin.Context) {
	// Получаем ID врача из URL
	doctorID, err := strconv.Atoi(c.Param("doctor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	// Получаем дату из query-параметров
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date parameter is required"})
		return
	}

	// Парсим дату
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// Получаем номер страницы (по умолчанию 1)
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be integer greater than 0"})
		return
	}

	// Вызываем сервис
	receptions, err := h.receptionService.GetReceptionsByDoctorAndDate(uint(doctorID), date, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}

