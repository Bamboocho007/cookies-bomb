package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	Email     string
}
