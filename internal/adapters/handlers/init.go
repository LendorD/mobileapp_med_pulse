package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"

	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handler struct {
	logger  *logging.Logger
	usecase interfaces.Usecases
	authUC  *usecases.AuthUsecase // Добавляем AuthUsecase напрямую
}

// // NewHandler создает новый экземпляр Handler со всеми зависимостями
// func NewHandler(usecase interfaces.Usecases) *Handler {
// 	return &Handler{
// 		usecase: usecase,
// 	}
// }

// NewHandler создает новый экземпляр Handler со всеми зависимостями
func NewHandler(usecase interfaces.Usecases, parentLogger *logging.Logger, authUC *usecases.AuthUsecase) *Handler {
	//logger := logging.NewLogger("HANDLER", "GENERAL", parentLogger)

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
func ProvideRouter(h *Handler) http.Handler {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(LoggingMiddleware(h.logger))

	// Роутеры авторизации
	authHandler := NewAuthHandler(h.authUC)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", gin.WrapF(authHandler.LoginDoctor))
	}

	// Основной маршрут
	baseRouter := r.Group("/api/v1")

	// Роутеры доктора
	doctorGroup := baseRouter.Group("/doctor")
	doctorGroup.POST("/", h.CreateDoctor)
	doctorGroup.GET("/:id", h.GetDoctorByID)
	doctorGroup.GET("/patients/:doc_id", h.GetPatientsByDoctorID)

	// Роутеры медкарты
	medCardGroup := baseRouter.Group("/medcard")
	medCardGroup.GET("/:id", h.GetMedCardByPatientID)
	medCardGroup.PUT("/:id", h.UpdateMedCard)

	receptionHospital := baseRouter.Group("/recepHospital")
	receptionHospital.GET("/:pat_id", h.GetReceptionsHospitalByPatientID)

	// Роутеры пациентов
	patientGroup := baseRouter.Group("/patients")
	patientGroup.POST("/", h.CreatePatient)
	patientGroup.GET("/:pat_id", h.GetPatientByID)
	patientGroup.DELETE("/:pat_id", h.DeletePatient)
	patientGroup.PATCH("/:pat_id", h.UpdatePatient)

	// Роутеры СМП
	// emergencyGroup := baseRouter.Group("/emergency-group")
	// emergencyGroup.GET("/:doctor_id", h.GetEmergencyReceptionsByDoctorAndDate)
	r.GET("/emergency/:doctor_id", h.GetEmergencyReceptionsByDoctorAndDate)

	// r.GET("/receptions/:doctor_id", h.GetReceptionsByDoctorAndDate)

	// Группа маршрутов для patientContactInfo
	contactInfoGroup := baseRouter.Group("/contact_info")
	contactInfoGroup.POST("/:pat_id", h.CreateContactInfo)
	contactInfoGroup.GET("/:pat_id", h.GetContactInfoByPatientID)

	return r
}
