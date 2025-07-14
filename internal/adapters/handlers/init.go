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

	// Группа аутентификации
	// Создаем AuthHandler

	authHandler := NewAuthHandler(h.authUC)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", gin.WrapF(authHandler.LoginDoctor))
	}

	baseRouter := r.Group("/api/v1")

	medCardGroup := baseRouter.Group("/medcard")
	medCardGroup.GET("/:pat_id", h.GetMedCardByPatientID)
	medCardGroup.PUT("/:pat_id", h.UpdateMedCard)

	// Группа маршрутов для patients
	receptionHospital := baseRouter.Group("/recepHospital")
	receptionHospital.GET("/:pat_id", h.GetReceptionsHospitalByPatientID)
	receptionHospital.PUT("/:recep_id", h.UpdateReceptionHospitalByReceptionID)

	// Роутеры пациентов
	patientGroup := baseRouter.Group("/patients")
	patientGroup.GET("/:doc_id", h.GetPatientsByDoctorID)
	patientGroup.GET("/recep_hosp/:pat_id", h.GetReceptionsHospitalByPatientID)

	// Временный для получения всех пациентов
	patientGroup.GET("/", h.GetAllPatients)

	// Роутеры СМП
	// emergencyGroup := baseRouter.Group("/emergency-group")
	// emergencyGroup.GET("/:doctor_id", h.GetEmergencyReceptionsByDoctorAndDate)
	r.GET("/emergency/:doctor_id", h.GetEmergencyReceptionsByDoctorAndDate)
	r.GET("/receps/:doctor_id", h.GetReceptionsByDoctorAndDate)

	return r
}
