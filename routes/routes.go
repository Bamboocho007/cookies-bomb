package routes

import "github.com/gofiber/fiber/v2"

func ApplyAllRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	apiV1.Post("/login", Login)
	apiV1.Post("/logout", Logout)
	apiV1.Post("/createUser", CreateUser)
}
