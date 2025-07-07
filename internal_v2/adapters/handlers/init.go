package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	logging "gitlab.com/devkit3/logger"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handler struct {
	logger  *logging.Logger
	usecase interfaces.Usecases
}

func NewHandler(parentLogger *logging.Logger, usecase interfaces.Usecases) *Handler {
	logger := logging.NewModuleLogger("HANDLER", "GENERAL", parentLogger)

	return &Handler{
		logger:  logger,
		usecase: usecase,
	}
}

func ProvideRouter(h *Handler) http.Handler {
	r := gin.Default()
	// r.Use(Logging(h.logger))
	// baseRouter := r.Group("/api/v1")

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа маршрутов для групп
	// bookingBookingGroup := baseRouter.Group("/booking-group")
	// bookingBookingGroup.POST("/", h.CreateGroup)
	// bookingBookingGroup.PUT("/", h.UpdateGroup)
	// bookingBookingGroup.GET("/:id", h.GetGroupByID)
	// bookingBookingGroup.GET("/", h.GetFilteredGroup)

	return r
}
