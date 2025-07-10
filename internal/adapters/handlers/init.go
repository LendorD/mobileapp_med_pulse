package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
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

func ProvideRouter(handlers *Handlers) http.Handler {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа аутентификации
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Auth.RegisterDoctor)
		authGroup.POST("/login", handlers.Auth.LoginDoctor)
	}

	return r
}
