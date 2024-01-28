package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthClaims struct {
	Id uuid.UUID `json:"id,omitempty"`
	jwt.RegisteredClaims
}
