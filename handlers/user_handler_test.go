package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"

	"web-api-in-go/handlers"
	"web-api-in-go/models"
)

type UserHandlerTestSuite struct {
	suite.Suite
	echo        *echo.Echo
	userHandler *handlers.UserHandler
	user        models.User
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.echo = echo.New()
	s.userHandler = handlers.NewUserHandler()

	s.user = models.User{
		ID:       1,
		Name:     "Leanne Graham",
		Username: "Bret",
		Email:    "Sincere@april.biz",
		Address: models.Address{
			Street:  "Kulas Light",
			Suite:   "Apt. 556",
			City:    "Gwenborough",
			Zipcode: "92998-3874",
			Geo: models.Geo{
				Lat: "-37.3159",
				Lng: "81.1496",
			},
		},
		Phone:   "1-770-736-8031 x56442",
		Website: "hildegard.org",
		Company: models.Company{
			Name:        "Romaguera-Crona",
			CatchPhrase: "Multi-layered client-server neural-net",
			Bs:          "harness real-time e-markets",
		},
	}
}

func (s *UserHandlerTestSuite) TestGetUser() {
	req, _ := http.NewRequest(echo.GET, "/users", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1") // DBを介さないため、固定値を指定

	s.userHandler.GetUser(c)
	s.Equal(http.StatusOK, rec.Code)

	var responseUser models.User
	s.NoError(json.Unmarshal(rec.Body.Bytes(), &responseUser))

	s.Equal(s.user, responseUser)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
