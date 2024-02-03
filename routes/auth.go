package routes

import (
	"net/http"

	"github.com/Bamboocho007/cookies-bomb/auth"
	"github.com/Bamboocho007/cookies-bomb/dto"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	loginDto := dto.LoginDto{}

	parseJsonError := c.Bind(&loginDto)

	if parseJsonError != nil {
		return parseJsonError
	}

	if validationError := loginDto.Validate(); validationError != nil {
		return validationError
	}

	jwtString, loginError := auth.Login(loginDto)

	if loginError != nil {
		return loginError
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_jwt"
	cookie.Value = jwtString

	c.SetCookie(cookie)

	return c.String(http.StatusOK, "Access allowed!")
}

func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "auth_jwt"
	cookie.Value = ""
	cookie.MaxAge = -1

	c.SetCookie(cookie)

	return nil
}

func CreateUser(c echo.Context) error {
	newUser := dto.NewUserDto{}
	parseBodyError := c.Bind(&newUser)
	if parseBodyError != nil {
		return parseBodyError
	}

	if validationError := newUser.Validate(); validationError != nil {
		return validationError
	}

	jwtString, createUserError := auth.CreateUser(newUser)
	if createUserError != nil {
		return createUserError
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_jwt"
	cookie.Value = jwtString

	c.SetCookie(cookie)

	return c.String(http.StatusOK, "User registered successfuly!")
}
