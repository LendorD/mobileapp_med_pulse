package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/swagger"

	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handler struct {
	logger  *logging.Logger
	usecase interfaces.Usecases
	authUC  *usecases.AuthUsecase
}

// NewHandler создает новый экземпляр Handler со всеми зависимостями
func NewHandler(usecase interfaces.Usecases, parentLogger *logging.Logger, authUC *usecases.AuthUsecase) *Handler {
	handlerLogger := parentLogger.WithPrefix("HANDLER")
	handlerLogger.Info("Handler initialized",
		"component", "GENERAL",
	)
	return &Handler{
		logger:  handlerLogger,
		usecase: usecase,
		authUC:  authUC,
	}
}

// ProvideRouter создает и настраивает маршруты
func ProvideRouter(h *Handler, swaggerCfg *swagger.Config) http.Handler {
	r := gin.Default()

	r.Use(LoggingMiddleware(h.logger))

	// Konkov: Инициализация Swagger, самого его вынес в middleware
	swagger.Setup(r, swaggerCfg)

	baseRouter := r.Group("/api/v1")

	doctorGroup := baseRouter.Group("/doctor")
	doctorGroup.POST("/", h.CreateDoctor)
	doctorGroup.GET("/:id", h.GetDoctorByID)

	medCardGroup := baseRouter.Group("/medcard")
	medCardGroup.GET("/:id", h.GetMedCardByPatientID)

	// Создаем AuthHandler
	authHandler := NewAuthHandler(h.authUC)

	// Группа аутентификации
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", gin.WrapF(authHandler.LoginDoctor))
	}

	// Группа маршрутов для patients
	patientGroup := baseRouter.Group("/patients")
	patientGroup.POST("/", h.CreatePatient)
	patientGroup.GET("/:pat_id", h.GetPatientByID)
	patientGroup.DELETE("/:pat_id", h.DeletePatient)
	patientGroup.PATCH("/:pat_id", h.UpdatePatient)

	// Группа маршрутов для patientContactInfo
	contactInfoGroup := baseRouter.Group("/contact_info")
	contactInfoGroup.POST("/:pat_id", h.CreateContactInfo)
	contactInfoGroup.GET("/:pat_id", h.GetContactInfoByPatientID)

	return r
}
