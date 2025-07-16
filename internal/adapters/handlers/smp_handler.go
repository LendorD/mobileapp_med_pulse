package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetReceptionsSMPByDoctorAndDate(c *gin.Context) {
	// Получаем doctor_id из URL
	doctorIDStr := c.Param("doc_id")
	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
		return
	}

	// Получаем номер страницы из query параметров
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page must be a positive integer"})
		return
	}

	// Получаем номер страницы из query параметров
	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "perPage must be a positive integer > 5"})
		return
	}

	// Вызываем usecase
	receptions, err := h.usecase.GetSMPReceptionsByEmergencyCall(uint(doctorID), page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}

func (h *Handler) GetReceptionWithMedServices(c *gin.Context) {
	// Парсинг ID
	id, err := strconv.ParseUint(c.Param("smp_id"), 10, 32)
	if err != nil {
		h.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'id' must be an integer", false)
		return
	}

	// Вызов usecase
	reception, err := h.usecase.GetReceptionWithMedServicesByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Reception not found",
			"code":  "reception_not_found",
		})
		return
	}

	c.JSON(http.StatusOK, reception)
}
