package config

import "os"

type EnvConfig struct {
	UserName     string
	UserPassword string
	DbName       string
	Host         string
	Port         string
}

func GetEnvConfig() *EnvConfig {
	return &EnvConfig{
		UserName:     getEnv("USER_ENV", ""),
		UserPassword: getEnv("USER_PASSWORD_ENV", ""),
		DbName:       getEnv("DB_NAME_ENV", ""),
		Host:         getEnv("HOST_ENV", ""),
		Port:         getEnv("PORT_ENV", ""),
	}
}

func getEnv(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
