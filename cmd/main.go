package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexanderMorozov1919/mobileapp/database"
)

func main() {
	// Инициализация БД с автомиграциями
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("✅ Database initialized and migrated successfully")

	// Простое тестовое использование
	// testDBOperations()

	// Ожидание сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

}
