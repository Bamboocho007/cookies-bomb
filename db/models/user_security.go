package models

import (
	"github.com/google/uuid"
)

type UserSecurity struct {
	IsConfirmed  byte      `db:"is_confirmed"`
	PasswordHash string    `db:"password_hash"`
	UserId       uuid.UUID `db:"user_id"`
}
