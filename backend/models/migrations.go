package models

import "github.com/Seinarukiro2/Secton/backend/config"

// AutoMigrate выполняет миграции
func AutoMigrate() {
	config.DB.AutoMigrate(&User{})
}
