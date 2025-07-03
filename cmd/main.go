package main

import (
	"github.com/AlexanderMorozov1919/mobileapp/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация БД с автомиграциями
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println(":) Database initialized and migrated successfully")

	// Простое тестовое использование
	// testDBOperations()

	// router = gin.Default()
	// router.GET("/smp/:doc_id", handlers.)

	// Ожидание сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

}
