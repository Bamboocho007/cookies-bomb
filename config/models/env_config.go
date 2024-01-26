package models

type EnvConfig struct {
	UserName     string
	UserPassword string
	DbName       string
	Host         string
	Port         string
	JwtSecret    string
	AppPort      string
}
