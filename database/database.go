package database

import (
	"log"
	"mobileapp/config"
	"mobileapp/internal/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	cfg := config.LoadDBConfig()

	db, err := config.InitDB(cfg)
	if err != nil {
		return err
	}

	DB = db

	// Автомиграция моделей
	err = migrateModels()
	if err != nil {
		return err
	}

	log.Println("Database tables migrated successfully")
	return nil
}

func migrateModels() error {
	return DB.AutoMigrate(
		&models.Doctor{},
		&models.Patient{},
		&models.ContactInfo{},
		&models.Reception{},
		&models.Allergy{},
	)
}

func GetDB() *gorm.DB {
	return DB
}
