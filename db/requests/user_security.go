package requests

import (
	"github.com/Bamboocho007/cookies-bomb/db"
	"github.com/Bamboocho007/cookies-bomb/db/models"
	"github.com/google/uuid"
)

func CreateUserSecurity(userSecurity models.UserSecurity) error {
	connectionError := db.PostgressConnection.Connect()

	if connectionError != nil {
		return connectionError
	}

	defer db.PostgressConnection.Close()

	_, execError := db.PostgressConnection.Db.NamedExec("INSERT INTO user_securities (user_id, password_hash, is_confirmed) VALUES(:user_id, :password_hash, :is_confirmed)", &userSecurity)

	if execError != nil {
		return execError
	}

	return nil
}

func GetUserSecurity(userId uuid.UUID, userSecurity *models.UserSecurity) error {
	connectionError := db.PostgressConnection.Connect()

	if connectionError != nil {
		return connectionError
	}

	defer db.PostgressConnection.Close()

	queryError := db.PostgressConnection.Db.Get(userSecurity, "SELECT user_id, password_hash, is_confirmed FROM user_securities WHERE user_id=$1", userId)

	if queryError != nil {
		return queryError
	}

	return nil
}
