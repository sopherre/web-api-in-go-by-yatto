package main

import (
	"web-api-in-go/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	e := echo.New()

	routes.SetupRoutes(e)
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	port := "8080"
	e.Start(":" + port)
}