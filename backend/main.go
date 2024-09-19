package main

import (
	"github.com/Seinarukiro2/Secton/backend/config"
	"github.com/Seinarukiro2/Secton/backend/models"
	"github.com/Seinarukiro2/Secton/backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Подключение к базе данных
	config.ConnectDB()

	// Выполняем миграции
	models.AutoMigrate()

	// Инициализация Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Настройка маршрутов
	routes.SetupRoutes(e)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
