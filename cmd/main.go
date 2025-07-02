package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Инициализация роутера
	r := gin.Default()

	// 2. Регистрация роутов (без логики БД)
	registerRoutes(r)

	// 3. Запуск сервера
	port := getEnv("APP_PORT", "8080")
	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func registerRoutes(r *gin.Engine) {
	// Заглушки для API
	r.POST("/auth/register", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "message": "Registration mock"})
	})

	r.POST("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "message": "Login mock", "token": "dummy_jwt_token"})
	})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
