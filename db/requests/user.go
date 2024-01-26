package requests

import (
	"github.com/Bamboocho007/cookies-bomb/db"
	"github.com/Bamboocho007/cookies-bomb/db/models"
	"github.com/google/uuid"
)

func CreateUser(user models.User) error {
	connectionError := db.PostgressConnection.Connect()

	if connectionError != nil {
		return connectionError
	}

	defer db.PostgressConnection.Close()

	_, execError := db.PostgressConnection.Db.NamedExec("INSERT INTO users (id, first_name, last_name, created_at, email) VALUES(:id, :first_name, :last_name, :created_at, :email)", &user)

	if execError != nil {
		return execError
	}

	return nil
}

func GetUserByEmail(email string, user *models.User) error {
	connectionError := db.PostgressConnection.Connect()

	if connectionError != nil {
		return connectionError
	}

	defer db.PostgressConnection.Close()

	queryError := db.PostgressConnection.Db.Get(user, "SELECT * FROM users WHERE email=$1", email)

	if queryError != nil {
		return queryError
	}

	return nil
}

func GetUserById(id uuid.UUID, user *models.User) error {
	connectionError := db.PostgressConnection.Connect()

	if connectionError != nil {
		return connectionError
	}

	defer db.PostgressConnection.Close()

	queryError := db.PostgressConnection.Db.Get(user, "SELECT (id, first_name, last_name, created_at, email) FROM users WHERE id=$1", id)

	if queryError != nil {
		return queryError
	}

	return nil
}
