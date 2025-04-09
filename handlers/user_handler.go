package handlers

import (
	"encoding/json"
	"net/http"

	"web-api-in-go/models"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUser godoc
// @Summary ユーザー情報の取得
// @Description 指定されたIDのユーザー情報を外部APIから取得します
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 200 {object} models.User
// @Failure 500 {string} string "Error fetching data or parsing"
// @Router /users/{id} [get]
func (u *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	url := "https://jsonplaceholder.typicode.com/users/" + id
	resp, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching data")
	}
	defer resp.Body.Close()

	var user models.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing data")
	}

	return c.JSON(http.StatusOK, user)
}
