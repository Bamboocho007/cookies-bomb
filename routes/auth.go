package routes

import (
	"encoding/json"

	"github.com/Bamboocho007/cookies-bomb/auth"
	"github.com/Bamboocho007/cookies-bomb/dto"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginDto := dto.LoginDto{}

	parseJsonError := json.Unmarshal(c.Body(), &loginDto)

	if parseJsonError != nil {
		return parseJsonError
	}

	jwtString, loginError := auth.Login(loginDto)

	if loginError != nil {
		return loginError
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_jwt",
		Value:    jwtString,
		HTTPOnly: true,
	})

	return c.SendString("Access allowed!")
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("auth_jwt")

	return nil
}

func CreateUser(c *fiber.Ctx) error {
	newUser := dto.NewUserDto{}
	parseBodyError := json.Unmarshal(c.Body(), &newUser)
	if parseBodyError != nil {
		return parseBodyError
	}

	jwtString, createUserError := auth.CreateUser(newUser)
	if createUserError != nil {
		return createUserError
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_jwt",
		Value:    jwtString,
		HTTPOnly: true,
	})

	return c.SendString("User registered successfuly!")
}
