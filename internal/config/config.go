package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port           string
	AllowedOrigins []string
}

type Config struct {
	App        AppConfig
	HTTPServer HTTPConfig
	Database   DatabaseConfig
	Redis      RedisConfig
	OneC       OneCConfig
	Logging    LoggerConfig
	Services   Services
	Server     ServerConfig // Добавляем ServerConfig в основную структуру
	JWTSecret  string
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Port: getEnv("SERVER_PORT", "8080"),
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhost:5173",
			"http://localhost:4200",
			"http://localhost:8081",
			"http://127.0.0.1:3000",
			"http://127.0.0.1:5173",
		},
	}
}

type AppConfig struct {
	Version string
}

type HTTPConfig struct {
	Port              string
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

type LoggerConfig struct {
	Enable     bool
	LogsDir    string
	Level      string
	Format     string
	SavingDays int
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type OneCConfig struct {
	BaseURL  string
	Timeout  time.Duration
	Username string
	Password string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Services struct {
	MobileApp Service
}

type Service struct {
	Host string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err.Error())
	}

	// Парсим timeout для 1С
	oneCTimeoutStr := getEnv("ONESC_TIMEOUT", "30s")
	oneCTimeout, err := time.ParseDuration(oneCTimeoutStr)
	if err != nil {
		return nil, fmt.Errorf("invalid ONESC_TIMEOUT: %w", err)
	}

	// Инициализируем Config с Server
	cfg := &Config{
		App: AppConfig{
			Version: getEnv("VERSION", "1.0.0"),
		},
		HTTPServer: HTTPConfig{
			Port:              getEnv("SERVER_PORT", "6004"),
			ReadTimeout:       time.Second * 10,
			ReadHeaderTimeout: time.Second * 20,
			WriteTimeout:      time.Second * 20,
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "192.168.29.138"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "db_name"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},

		OneC: OneCConfig{
			BaseURL:  getEnv("ONESC_BASE_URL", "http://localhost:8081/hs/api"),
			Timeout:  oneCTimeout,
			Username: getEnv("ONESC_USERNAME", ""),
			Password: getEnv("ONESC_PASSWORD", ""),
		},
		Logging: LoggerConfig{
			Enable:     getEnvAsBool("LOGGER_ENABLE", true),
			LogsDir:    getEnv("LOGGER_LOGS_DIR", "./logs"),
			Level:      getEnv("LOGGER_LOG_LEVEL", "DEBUG"),
			SavingDays: getEnvAsInt("LOGGER_SAVING_DAYS", 5),
		},
		Services: Services{
			MobileApp: Service{
				Host: getEnv("API_URL", "http://localhost:8080"),
			},
		},
		Server: ServerConfig{ // Явно инициализируем Server
			Port: getEnv("SERVER_PORT", "8080"),
			AllowedOrigins: []string{
				"http://localhost:3000",
				"http://localhost:5173",
				"http://localhost:4200",
				"http://localhost:8081",
				"http://127.0.0.1:3000",
				"http://127.0.0.1:5173",
			},
		},
	}

	cfg.JWTSecret = getEnv("JWT_SECRET", "")
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	val, _ := strconv.ParseBool(value)
	return val
}
