package main

import (
	"fmt"
	"time"

	"github.com/Bamboocho007/cookies-bomb/config"
	"github.com/Bamboocho007/cookies-bomb/db"
	"github.com/Bamboocho007/cookies-bomb/db/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	envConfig := config.GetEnvConfig()
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envConfig.UserName, envConfig.UserPassword, envConfig.Host, envConfig.Port, envConfig.DbName)

	store := db.NewPostgresStore(databaseUrl)

	dbConnection := store.Connect()

	if dbConnection != nil {
		panic(dbConnection.Error())
	}

	defer store.Db.Close()

	var users []models.User

	users = append(users, models.User{
		Id:        uuid.New(),
		FirstName: "User 1",
		LastName:  "User last name 1",
		CreatedAt: time.Now().UTC(),
		Email:     "testmail1@gmal.com",
	})

	users = append(users, models.User{
		Id:        uuid.New(),
		FirstName: "User 2",
		LastName:  "User last name 2",
		CreatedAt: time.Now().UTC(),
		Email:     "testmail2@gmal.com",
	})

	users = append(users, models.User{
		Id:        uuid.New(),
		FirstName: "User 3",
		LastName:  "User last name 3",
		CreatedAt: time.Now().UTC(),
		Email:     "testmail3@gmal.com",
	})

	tx := store.Db.MustBegin()

	for _, user := range users {
		tx.MustExec(
			"INSERT INTO users (id, first_name, last_name, created_at, email) VALUES ($1, $2, $3, $4, $5)",
			user.Id, user.FirstName, user.LastName, user.CreatedAt, user.Email)
	}

	tx.Commit()

	fmt.Println("Project started!")
}
