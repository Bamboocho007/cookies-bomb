package models

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	Email string
	jwt.RegisteredClaims
}