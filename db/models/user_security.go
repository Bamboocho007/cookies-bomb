package models

import "github.com/google/uuid"

type UserCecurity struct {
	IsConfirmed bool `db:"is_confirmed"`
	Password    string
	UserId      uuid.UUID `db:"user_id"`
}
