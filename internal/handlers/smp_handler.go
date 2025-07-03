package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/gin-gonic/gin"
)

type SmpHandlers struct {
	smpService services.SmpService
}

func NewOrderHandler(smpService services.SmpService) *SmpHandlers {
	return &SmpHandlers{
		smpService: smpService,
	}
}

func (h *SmpHandlers) GetAllAmbulanceCallings(c *gin.Context) {
	doctorIDStr := c.Param("id")
	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 32)
	if err != nil {
		// Обработка ошибки (например, возврат статуса 400 Bad Request)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID format"})
		return
	}

	// Преобразуем uint64 в uint (32-битное)
	doctorIDUint := uint(doctorID)
	callings, err := h.smpService.GetCallings(doctorIDUint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order request"})
		return
	}
	c.JSON(http.StatusOK, callings)
}
