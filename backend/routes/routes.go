package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/Seinarukiro2/Secton/backend/controllers"
)

// SetupRoutes инициализирует маршруты для приложения
func SetupRoutes(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)
}
