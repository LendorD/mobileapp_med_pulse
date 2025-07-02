package main

import (
	"log"
	"mobileapp/database"
	"mobileapp/internal/models"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация БД с автомиграциями
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("✅ Database initialized and migrated successfully")

	// Простое тестовое использование
	testDBOperations()

	// Ожидание сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
}

func testDBOperations() {
	db := database.GetDB()

	// Пример создания доктора
	doctor := models.Doctor{
		FirstName:      "Иван",
		Surname:        "Петров",
		Login:          "ivan.petrov",
		PasswordHash:   "securehash",
		Specialization: "Кардиолог",
	}

	if err := db.Create(&doctor).Error; err != nil {
		log.Printf("Error creating doctor: %v", err)
	} else {
		log.Printf("Created doctor with ID: %d", doctor.ID)
	}
}
