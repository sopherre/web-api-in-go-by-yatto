package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck godoc
// @Summary ヘルスチェック
// @Description サーバーの状態確認用
// @Tags health
// @Success 200 {string} string "OK!"
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")
}

// HelloWorld godoc
// @Summary Hello World
// @Description Go API の動作確認用
// @Tags hello
// @Success 200 {string} string "Hello, Go World!"
// @Router /hello [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Go World!")
}
