package repositories

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSQLiteDB(cfg *config.Config) (*gorm.DB, error) {
	dir := filepath.Dir(cfg.Database.LocalDBPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать директорию для local.db: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.Database.LocalDBPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к SQLite (local): %w", err)
	}

	return db, nil
}
