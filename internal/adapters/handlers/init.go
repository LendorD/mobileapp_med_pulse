package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"github.com/gin-contrib/cors"
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
	service interfaces.Service
	authUC  *usecases.AuthUsecase // Добавляем AuthUsecase напрямую
}

// NewHandler создает новый экземпляр Handler со всеми зависимостями
func NewHandler(usecase interfaces.Usecases, parentLogger *logging.Logger, service interfaces.Service, authUC *usecases.AuthUsecase) *Handler {
	handlerLogger := parentLogger.WithPrefix("HANDLER")
	handlerLogger.Info("Handler initialized",
		"component", "GENERAL",
	)
	return &Handler{
		logger:  handlerLogger,
		usecase: usecase,
		service: service,
		authUC:  authUC,
	}
}

// ProvideRouter создает и настраивает маршруты
func ProvideRouter(h *Handler, cfg *config.Config) http.Handler {
	r := gin.Default()

	// В ProvideRouter:
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(LoggingMiddleware(h.logger))

	// Группа аутентификации
	authHandler := NewAuthHandler(h.authUC)
	authGroup := r.Group("/auth")
	authGroup.POST("/login", gin.WrapF(authHandler.LoginDoctor))

	baseRouter := r.Group("/api/v1")

	// Группа медкарты пациента
	medCardGroup := baseRouter.Group("/medcard")
	medCardGroup.GET("/:pat_id", h.GetMedCardByPatientID)
	medCardGroup.PUT("/:pat_id", h.UpdateMedCard)
	// TODO: Дописать добавление новых аллергий пациента и их изменение

	r.GET("/receps/:doctor_id", h.GetReceptionsSMPByDoctorAndDate)

	// Группа маршрутов для заключений
	receptionHospital := baseRouter.Group("/recepHospital")
	receptionHospital.GET("/:doc_id", h.GetReceptionsHospitalByDoctorAndDate)
	receptionHospital.GET("/patients/:pat_id", h.GetReceptionsHospitalByPatientID)
	receptionHospital.PUT("/:recep_id", h.UpdateReceptionHospitalByReceptionID)

	// Роутеры пациентов
	patientGroup := baseRouter.Group("/patients")
	patientGroup.GET("/:doc_id", h.GetPatientsByDoctorID)
	patientGroup.GET("/recep_hosp/:pat_id", h.GetReceptionsHospitalByPatientID)

	// Временный для получения всех пациентов
	patientGroup.GET("/", h.GetAllPatients)

	// Роутеры СМП
	emergencyGroup := baseRouter.Group("/emergencyGroup")

	emergencyGroup.GET("/:doc_id", h.GetEmergencyCallsByDoctorAndDate)
	emergencyGroup.GET("/:doc_id/smps", h.GetReceptionsSMPByDoctorAndDate)
	emergencyGroup.GET("/:doc_id/smps/:smp_id", h.GetReceptionWithMedServices)

	return r
}
