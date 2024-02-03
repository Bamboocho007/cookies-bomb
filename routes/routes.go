package routes

import (
	"github.com/labstack/echo/v4"
)

func ApplyAllRoutes(app *echo.Echo) {
	apiV1 := app.Group("/api/v1")
	apiV1.POST("/login", Login)
	apiV1.POST("/logout", Logout)
	apiV1.POST("/createUser", CreateUser)
}
