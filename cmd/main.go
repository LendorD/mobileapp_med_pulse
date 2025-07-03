package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/database"
	"github.com/AlexanderMorozov1919/mobileapp/internal/handlers"
	"github.com/AlexanderMorozov1919/mobileapp/internal/repository"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация БД
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Инициализация зависимостей авторизации
	authRepo := repository.NewAuthRepository(database.GetDB())
	recepRepo := repository.NewReceptionRepository(database.GetDB())
	authService := services.NewAuthService(
		authRepo,
		"your_jwt_secret_key", // Замените на реальный секретный ключ
		24*time.Hour,          // Время жизни токена
	)
	recepService := services.NewReceptionService(recepRepo)

	authHandler := handlers.NewAuthHandler(authService)
	recepHandler := handlers.NewReceptionHandler(recepService)

	// Настройка роутера
	router := gin.Default()

	// Роуты авторизации
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.POST("/newRecep", recepHandler.CreateReception)
	router.GET("/main/:doctor_id/receps", recepHandler.GetDoctorReceptions)

	// Защищенные роуты (пример)
	// authorized := router.Group("/")
	// authorized.Use(middleware.AuthMiddleware(authService))
	// {
	//     authorized.GET("/profile", profileHandler)
	// }

	// Запуск сервера
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Println("Server started on :8080")

	// Ожидание сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
}
