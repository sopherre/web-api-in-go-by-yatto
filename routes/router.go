package routes

import (
	"net/http"
	"web-api-in-go/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go World!")
	})

	e.GET("/users/:id", handlers.GetUser)

	e.GET("/tasks", handlers.GetTasks)
	e.GET("/tasks/:id", handlers.GetTask)
	e.POST("/tasks", handlers.CreateTask)
	e.PUT("/tasks/:id", handlers.UpdateTask)
	e.DELETE("/tasks/:id", handlers.DeleteTask)
}