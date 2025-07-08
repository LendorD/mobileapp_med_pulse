package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handler struct {
	// logger  *logging.Logger
	usecase interfaces.Usecases
}

func NewHandler(usecase interfaces.Usecases) *Handler {
	//logger := logging.NewLogger("HANDLER", "GENERAL", parentLogger)

	return &Handler{
		//logger:  logger,
		usecase: usecase,
	}
}

func ProvideRouter(h *Handler) http.Handler {
	r := gin.Default()
	// r.Use(Logging(h.logger))
	// baseRouter := r.Group("/api/v1")

	r.GET("/main/:doctor_id", h.GetReceptionsByDoctorAndDate)
	r.GET("/doctors/:doctor_id/receptions", h.GetReceptionsByDoctorAndDate) // Новый маршрут из Swagger
	// Группа маршрутов для групп
	// bookingBookingGroup := baseRouter.Group("/booking-group")
	// bookingBookingGroup.POST("/", h.CreateGroup)
	// bookingBookingGroup.PUT("/", h.UpdateGroup)
	// bookingBookingGroup.GET("/:id", h.GetGroupByID)
	// bookingBookingGroup.GET("/", h.GetFilteredGroup)

	return r
}
