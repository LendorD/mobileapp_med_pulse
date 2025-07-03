package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/gin-gonic/gin"
)

type ReceptionHandler struct {
	service services.ReceptionService
}

func NewReceptionHandler(service services.ReceptionService) *ReceptionHandler {
	return &ReceptionHandler{service: service}
}

// CreateReception creates a new reception entry
// @Summary Create a new reception
// @Description Add a new patient reception record
// @Tags receptions
// @Accept json
// @Produce json
// @Param reception body models.Reception true "Reception details"
// @Success 201 {object} models.Reception
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /receptions [post]
func (h *ReceptionHandler) CreateReception(c *gin.Context) {
	var reception models.Reception
	if err := c.ShouldBindJSON(&reception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.CreateReception(&reception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reception)
}

// GetReception retrieves a reception by ID
// @Summary Get reception by ID
// @Description Get a specific reception by its ID
// @Tags receptions
// @Produce json
// @Param id path int true "Reception ID"
// @Success 200 {object} models.Reception
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /receptions/{id} [get]
func (h *ReceptionHandler) GetReception(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reception ID"})
		return
	}

	reception, err := h.service.GetReceptionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reception not found"})
		return
	}

	c.JSON(http.StatusOK, reception)
}

// UpdateReception updates an existing reception
// @Summary Update reception
// @Description Update an existing reception record
// @Tags receptions
// @Accept json
// @Produce json
// @Param id path int true "Reception ID"
// @Param reception body models.Reception true "Updated reception details"
// @Success 200 {object} models.Reception
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /receptions/{id} [put]
func (h *ReceptionHandler) UpdateReception(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reception ID"})
		return
	}

	var reception models.Reception
	if err := c.ShouldBindJSON(&reception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	reception.ID = uint(id)
	if err := h.service.UpdateReception(&reception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reception)
}

// CancelReception cancels a reception
// @Summary Cancel reception
// @Description Cancel an existing reception
// @Tags receptions
// @Accept json
// @Produce json
// @Param id path int true "Reception ID"
// @Param reason body CancelRequest true "Cancellation reason"
// @Success 200 {object} models.Reception
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /receptions/{id}/cancel [post]
func (h *ReceptionHandler) CancelReception(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reception ID"})
		return
	}

	var request struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.CancelReception(uint(id), request.Reason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reception, err := h.service.GetReceptionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reception not found"})
		return
	}

	c.JSON(http.StatusOK, reception)
}

// GetDoctorReceptions retrieves receptions by doctor
// @Summary Get receptions by doctor
// @Description Get all receptions for a specific doctor
// @Tags receptions
// @Produce json
// @Param doctor_id path int true "Doctor ID"
// @Param date query string false "Filter by date (YYYY-MM-DD)"
// @Success 200 {array} models.Reception
// @Failure 400 {object} ErrorResponse
// @Router /doctors/{doctor_id}/receptions [get]
func (h *ReceptionHandler) GetDoctorReceptions(c *gin.Context) {
	doctorID, err := strconv.Atoi(c.Param("doctor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var date *time.Time
	if dateParam := c.Query("date"); dateParam != "" {
		parsedDate, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		date = &parsedDate
	}

	receptions, err := h.service.GetDoctorReceptions(uint(doctorID), date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}

// GetPatientReceptions retrieves receptions by patient
// @Summary Get receptions by patient
// @Description Get all receptions for a specific patient
// @Tags receptions
// @Produce json
// @Param patient_id path int true "Patient ID"
// @Success 200 {array} models.Reception
// @Failure 400 {object} ErrorResponse
// @Router /patients/{patient_id}/receptions [get]
func (h *ReceptionHandler) GetPatientReceptions(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	receptions, err := h.service.GetPatientReceptions(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}

// CompleteReception marks a reception as completed
// @Summary Complete reception
// @Description Mark a reception as completed with diagnosis
// @Tags receptions
// @Accept json
// @Produce json
// @Param id path int true "Reception ID"
// @Param data body CompleteRequest true "Completion data"
// @Success 200 {object} models.Reception
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /receptions/{id}/complete [post]
func (h *ReceptionHandler) CompleteReception(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reception ID"})
		return
	}

	var request struct {
		Diagnosis       string `json:"diagnosis" binding:"required"`
		Recommendations string `json:"recommendations"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.CompleteReception(uint(id), request.Diagnosis, request.Recommendations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reception, err := h.service.GetReceptionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reception not found"})
		return
	}

	c.JSON(http.StatusOK, reception)
}
