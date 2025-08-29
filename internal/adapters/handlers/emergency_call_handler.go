package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetEmergencyCallsByDoctorAndDate godoc
// @Summary Получить экстренные приёмы врача по дате
// @Description Возвращает список экстренных приёмов, назначенных врачу на указанную дату, с пагинацией
// @Tags SMP
// @Accept json
// @Produce json
// @Param doc_id path uint true "ID врача"
// @Param date query string false "Дата в формате YYYY-MM-DD"
// @Param page query int false "Номер страницы" default(1)
// @Param perPage query int false "Количество записей на страницу" default(5)
// @Success 200 {array} entities.EmergencyCall "Список приёмов"
// @Failure 400 {object} IncorrectFormatError "Некорректный запрос"
// @Failure 500 {object} InternalServerError "Внутренняя ошибка"
// @Router /emergency/{doc_id} [get]
func (h *Handler) GetEmergencyCallsByDoctorAndDate(c *gin.Context) {
	// Получаем ID врача
	doctorID, err := strconv.ParseUint(c.Param("doc_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	// Получаем дату из query параметров
	dateStr := c.Query("date")
	var date time.Time
	if dateStr != "" {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
			return
		}
	} else {
		// Если дата не указана, используем текущую дату
		date = time.Now()
	}

	// Получаем номер страницы
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be integer greater than 0"})
		return
	}

	perPageStr := c.DefaultQuery("perPage", "10")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be integer greater than 0"})
		return
	}

	// Вызываем usecase
	receptions, err := h.usecase.GetEmergencyCallsByDoctorAndDate(uint(doctorID), date, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}

// CloseEmergencyCall godoc
// @Summary Закрыть экстренный вызов
// @Description Возвращает экстренный вызов
// @Tags Calls
// @Accept json
// @Produce json
// @Param call_id path uint true "ID emergencyCall"
// @Success 200 {array} entities.EmergencyCall "Экстренный вызов"
// @Failure 400 {object} IncorrectFormatError "Некорректный запрос"
// @Failure 500 {object} InternalServerError "Внутренняя ошибка"
// @Router /emergency/{call_id} [patch]
func (h *Handler) CloseEmergencyCall(c *gin.Context) {
	callID, err := strconv.ParseUint(c.Param("call_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid call ID"})
		return
	}
	receptions, err := h.usecase.CloseEmergencyCall(uint(callID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}

func (h *Handler) GetPdf(c *gin.Context) {
	fmt.Println("Зашли для получения pdf")
	dir, err := os.Getwd() // текущая рабочая директория, обычно это cmd/app
	fmt.Println("директория:", dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get working directory"})
		return
	}

	pdfPath := filepath.Join(dir, "assets", "mobileapp.pdf")

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "PDF not found"})
		return
	}

	c.File(pdfPath) // Gin сам установит нужные заголовки
}
