package auth

import (
	"time"

	"github.com/Bamboocho007/cookies-bomb/db/models"
	"github.com/Bamboocho007/cookies-bomb/db/requests"
	"github.com/Bamboocho007/cookies-bomb/dto"
	"github.com/google/uuid"
)

func Login(loginDto dto.LoginDto) (string, error) {
	var user models.User
	getUserError := requests.GetUserByEmail(loginDto.Email, &user)
	if getUserError != nil {
		return "", getUserError
	}

	var userSecurity models.UserSecurity
	getUserSecurityError := requests.GetUserSecurity(user.Id, &userSecurity)
	if getUserSecurityError != nil {
		return "", getUserSecurityError
	}

	compareHashError := CompareHash(userSecurity.PasswordHash, loginDto.Password)
	if compareHashError != nil {
		return "", compareHashError
	}

	jwtString, generateJWTError := GenerateJWT(user.Id)
	if generateJWTError != nil {
		return "", generateJWTError
	}

	return jwtString, nil
}

func CreateUser(newUser dto.NewUserDto) (string, error) {
	newUserId := uuid.New()
	passwordHash, hashPasswordError := HashPassword(newUser.Password)
	if hashPasswordError != nil {
		return "", hashPasswordError
	}

	userToSave := models.User{
		Id:        newUserId,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		CreatedAt: time.Now().UTC(),
		Email:     newUser.Email,
	}
	userSecurityToSave := models.UserSecurity{
		UserId:       newUserId,
		PasswordHash: passwordHash,
		IsConfirmed:  0,
	}

	createUserError := requests.CreateUser(userToSave)
	if createUserError != nil {
		return "", createUserError
	}

	createUserSecurityError := requests.CreateUserSecurity(userSecurityToSave)
	if createUserSecurityError != nil {
		return "", createUserSecurityError
	}

	return GenerateJWT(userToSave.Id)
}
