package handlers

import (
	"encoding/json"
	"net/http"

	"web-api-in-go/models"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
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
