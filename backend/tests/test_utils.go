package tests

import (
	"log"
	"github.com/Seinarukiro2/Secton/backend/config"
	"github.com/Seinarukiro2/Secton/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitTestDB инициализирует базу данных для тестов
func InitTestDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Выполняем миграцию моделей
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Устанавливаем глобальную переменную DB в config
	config.DB = db
}
