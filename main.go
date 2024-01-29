package main

import (
	"fmt"

	"github.com/Bamboocho007/cookies-bomb/common/models"
	"github.com/Bamboocho007/cookies-bomb/config"
	"github.com/Bamboocho007/cookies-bomb/db"
	"github.com/Bamboocho007/cookies-bomb/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.LoadEnvConfig()
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.LoadedEnvConfig.UserName, config.LoadedEnvConfig.UserPassword, config.LoadedEnvConfig.Host, config.LoadedEnvConfig.Port, config.LoadedEnvConfig.DbName)
	db.InitPostgresStore(databaseUrl)

	engine := html.New("./layouts", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			if errorResponse, ok := err.(*models.ErrorResponse); ok {
				ctx.Status(fiber.ErrBadRequest.Code).JSON(errorResponse)
				return nil
			}

			return err
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Hello": "Hello world!",
		}, "menu", "footer", "header", "base")
	})

	routes.ApplyAllRoutes(app)

	app.Listen(fmt.Sprintf(":%s", config.LoadedEnvConfig.AppPort))

	fmt.Println("Project started!")
}
