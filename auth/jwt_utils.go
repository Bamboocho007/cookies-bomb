package auth

import (
	"errors"
	"time"

	"github.com/Bamboocho007/cookies-bomb/auth/models"
	"github.com/Bamboocho007/cookies-bomb/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userEmail string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &models.AuthClaims{
		Email: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.LoadedEnvConfig.JwtSecret))
}

func VerifyJWT(jwtString string) (*jwt.Token, error) {
	// use claims in the futore to verify token deeper
	claims := &models.AuthClaims{}
	token, parseError := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {

		return config.LoadedEnvConfig.JwtSecret, nil
	})

	if parseError != nil {
		return nil, parseError
	}

	if !token.Valid {
		return nil, errors.New("token not valid")
	}

	return token, nil
}

func RefreshToken(jwtString string) (string, error) {
	claims := &models.AuthClaims{}
	token, parseError := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {

		return config.LoadedEnvConfig.JwtSecret, nil
	})

	if parseError != nil {
		return "", parseError
	}

	if !token.Valid {
		return "", errors.New("token not valid")
	}

	if time.Until(claims.IssuedAt.Time) > time.Second*30 {
		return "", errors.New("token is fresh enough")
	}

	return GenerateJWT(claims.Email)
}
