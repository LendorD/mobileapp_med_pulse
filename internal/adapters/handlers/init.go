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
	baseRouter := r.Group("/api/v1")

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа маршрутов для доктора
	doctorGroup := baseRouter.Group("/doctor-group")
	doctorGroup.POST("/", h.CreateDoctor)
	doctorGroup.GET("/:id", h.GetDoctorByID)

	return r

}
