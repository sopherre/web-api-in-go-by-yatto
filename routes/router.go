package routes

import (
	"net/http"
	"web-api-in-go/db"
	"web-api-in-go/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(db *db.Database, e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go World!")
	})

	e.GET("/users/:id", handlers.GetUser)

	handlers := handlers.NewTaskHandler(db)
	handlers.SetupRoutes(e)
}
