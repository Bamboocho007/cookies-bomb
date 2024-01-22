package config

import (
	"os"

	"github.com/Bamboocho007/cookies-bomb/config/constants"
	"github.com/Bamboocho007/cookies-bomb/config/models"
	"github.com/joho/godotenv"
)

var LoadedEnvConfig *models.EnvConfig

func LoadEnvConfig() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	LoadedEnvConfig = &models.EnvConfig{
		UserName:     getEnv(constants.UserEnv, ""),
		UserPassword: getEnv(constants.UserPasswordEnv, ""),
		DbName:       getEnv(constants.DbNameEnv, ""),
		Host:         getEnv(constants.HostEnv, ""),
		Port:         getEnv(constants.PortEnv, ""),
		JwtSecret:    getEnv(constants.JwtSecretEnv, ""),
	}
}

func getEnv(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
