package main

import (
	"fmt"

	"github.com/Bamboocho007/cookies-bomb/common/models"
	"github.com/Bamboocho007/cookies-bomb/config"
	"github.com/Bamboocho007/cookies-bomb/db"
	"github.com/Bamboocho007/cookies-bomb/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnvConfig()
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.LoadedEnvConfig.UserName, config.LoadedEnvConfig.UserPassword, config.LoadedEnvConfig.Host, config.LoadedEnvConfig.Port, config.LoadedEnvConfig.DbName)
	db.InitPostgresStore(databaseUrl)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			if errorResponse, ok := err.(*models.ErrorResponse); ok {
				ctx.Status(fiber.ErrBadRequest.Code).JSON(errorResponse)
				return nil
			}

			return err
		},
	})

	routes.ApplyAllRoutes(app)

	app.Listen(fmt.Sprintf(":%s", config.LoadedEnvConfig.AppPort))

	fmt.Println("Project started!")
}
