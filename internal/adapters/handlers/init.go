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

	// Глобальные middleware
	// r.Use(Logging(h.logger))

	// Документация Swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 группа
	api := r.Group("/api/v1")
	{
		// Группа маршрутов для пациентов
		patients := api.Group("/patients")
		{
			patients.POST("", h.CreatePatient)       // POST /api/v1/patients
			patients.PUT("/:id", h.UpdatePatient)    // PUT /api/v1/patients/:id
			patients.GET("/:id", h.GetPatientByID)   // GET /api/v1/patients/:id
			patients.DELETE("/:id", h.DeletePatient) // DELETE /api/v1/patients/:id
		}

		// Другие группы маршрутов можно добавить аналогично
		// booking := api.Group("/booking")
		// {
		//     booking.POST("", h.CreateBooking)
		// }
	}

	return r
}
