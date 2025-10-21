// @title ClinicHub API
// @version 1.0.0
// @description API для работы с приёмами пациентов
// @contact.name API Support
// @contact.email support@example.com
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer <JWT>" to authenticate
package main

import (
	_ "github.com/AlexanderMorozov1919/mobileapp/docs"
	"github.com/AlexanderMorozov1919/mobileapp/internal/app"
)

func main() {
	app.New().Run()
}
