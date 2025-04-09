package routes

import (
	"web-api-in-go/db"
	"web-api-in-go/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(db *db.Database, e *echo.Echo) {
	e.GET("/health", handlers.HealthCheck)
	e.GET("/hello", handlers.HelloWorld)

	userHandlers := handlers.NewUserHandler()
	e.GET("/users/:id", userHandlers.GetUser)

	taskHandlers := handlers.NewTaskHandler(db)
	taskHandlers.SetupRoutes(e)
}
