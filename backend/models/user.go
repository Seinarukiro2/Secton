package models

import "gorm.io/gorm"

// User представляет пользователя с Telegram ID
type User struct {
	gorm.Model
	TelegramID string `json:"telegram_id" gorm:"unique;not null"`
}
