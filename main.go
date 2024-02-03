package main

import (
	"fmt"
	"net/http"

	"github.com/Bamboocho007/cookies-bomb/common/models"
	"github.com/Bamboocho007/cookies-bomb/config"
	"github.com/Bamboocho007/cookies-bomb/db"
	"github.com/Bamboocho007/cookies-bomb/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnvConfig()
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.LoadedEnvConfig.UserName, config.LoadedEnvConfig.UserPassword, config.LoadedEnvConfig.Host, config.LoadedEnvConfig.Port, config.LoadedEnvConfig.DbName)
	db.InitPostgresStore(databaseUrl)

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError

		if errorResponse, ok := err.(*models.ErrorResponse); ok {
			c.Logger().Error(errorResponse)
			return
		}

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		c.Logger().Error(err)

		errorPage := fmt.Sprintf("%d.html", code)
		if err := c.File(errorPage); err != nil {
			c.Logger().Error(err)
		}
	}

	routes.ApplyAllRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.LoadedEnvConfig.AppPort)))
}
