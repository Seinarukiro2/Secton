package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/Seinarukiro2/Secton/backend/models"
	"github.com/Seinarukiro2/Secton/backend/config"
)

// CreateUser создает нового пользователя на основе Telegram ID
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Асинхронная запись пользователя в базу данных
	go func() {
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
		}
	}()

	return c.JSON(http.StatusCreated, user)
}
