package main

import (
	"web-api-in-go/db"
	"web-api-in-go/routes"

	_ "web-api-in-go/docs" // ← Swagger docs を読み込むために必要（swag init で生成される）

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // Swagger UI 用
)

// @title Task Management API
// @version 1.0
// @description Go + Echo を使ったシンプルなタスク管理API
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	// DB 初期化（Gorm）
	dbInst := db.Init()
	defer dbInst.Close()

	e := echo.New()

	// ルーティング設定（handlers内でまとめて管理）
	routes.SetupRoutes(dbInst, e)

	// ミドルウェア：ログ出力とリカバリ（パニック対策）
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Swagger ドキュメントへのルート追加（例: /swagger/index.html）
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// ポート指定して起動
	port := "8080"
	e.Logger.Fatal(e.Start(":" + port))
}
