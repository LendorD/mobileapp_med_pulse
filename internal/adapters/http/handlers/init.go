package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	middleware "github.com/AlexanderMorozov1919/mobileapp/internal/middleware/jwt"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/swagger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/swaggo/files"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handler struct {
	logger  *logging.Logger
	usecase interfaces.Usecases
	service interfaces.Service
}

// NewHandler создает новый экземпляр Handler со всеми зависимостями
func NewHandler(usecase interfaces.Usecases, parentLogger *logging.Logger, service interfaces.Service) *Handler {
	handlerLogger := parentLogger.WithPrefix("HANDLER")
	handlerLogger.Info("Handler initialized",
		"component", "GENERAL",
	)
	return &Handler{
		logger:  handlerLogger,
		usecase: usecase,
		service: service,
	}
}

// ProvideRouter создает и настраивает маршруты
func ProvideRouter(h *Handler, ws *WebsocketHandler, cfg *config.Config, swagCfg *swagger.Config) http.Handler {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger-роутер
	swagger.Setup(r, swagCfg)

	// Logger
	r.Use(LoggingMiddleware(h.logger))

	// Общая группа для API
	baseRouter := r.Group("/api/v1")

	protected := baseRouter.Group("/")
	protected.Use(middleware.JWTAuth(cfg.JWTSecret))

	//Версия
	baseRouter.GET("/version", h.GetVersionProject)

	// Авторизация
	authGroup := baseRouter.Group("/auth")
	authGroup.POST("/", h.LoginDoctor)

	// WebSocket-группа
	wsGroup := r.Group("/ws/notification")
	// wsGroup.Use(middleware.JWTAuth(cfg.JWTSecret))

	wsGroup.GET("/register/:user_id", ws.Register)
	wsGroup.GET("/unregister/:user_id", ws.Unregister)

	//Запросы от 1С
	webhook := protected.Group("webhook")
	webhook.POST("/onec/receptions", h.OneCWebhook)          // Получение заявок
	webhook.POST("/onec/patients", h.OneCPatientListWebhook) // получение списка пациентов
	webhook.POST("/onec/auth", h.OneCAuthWebhook)            // Получение списка авторизации

	// Пациенты
	patientGroup := protected.Group("/patients")
	patientGroup.GET("", h.GetPatientList) // Отдаёт список всех пациентов c пагинацией (1С)

	// Медкарты (Больше не формируется а получаются от 1С)
	medCardGroup := protected.Group("/medcard")
	medCardGroup.GET("/:pat_id", h.GetMedCardByPatientID)
	medCardGroup.PUT("/:pat_id", h.UpdateMedCard)

	// Выезд
	emergencyGroup := protected.Group("/emergency")

	//Подписи пациентов
	emergencyGroup.GET("/signature/:recep_id", h.GetSignature)
	emergencyGroup.POST("/signature/:recep_id", h.SaveSignature)

	// emergencyGroup.POST("/pdf/:rec_id", h.UploadPdf)

	return r
}
