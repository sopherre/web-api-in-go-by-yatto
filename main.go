package main

import (
	"web-api-in-go/db"
	"web-api-in-go/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbInst := db.Init()
	defer dbInst.Close()

	e := echo.New()

	routes.SetupRoutes(dbInst, e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	port := "8080"
	e.Start(":" + port)
}
