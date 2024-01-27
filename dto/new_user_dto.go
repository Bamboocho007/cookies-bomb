package dto

import (
	"fmt"

	"github.com/Bamboocho007/cookies-bomb/common/models"
	"github.com/go-playground/validator/v10"
)

type NewUserDto struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

func (newUser *NewUserDto) Validate() *models.ErrorResponse {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validationErrors := validate.Struct(newUser)

	if validationErrors != nil {
		if err, ok := validationErrors.(*validator.InvalidValidationError); ok {
			return &models.ErrorResponse{Message: err.Error()}
		}

		responseErrors := &models.ErrorResponse{Message: "Body Validation error"}

		for _, err := range validationErrors.(validator.ValidationErrors) {
			responseErrors.Errors = append(responseErrors.Errors, models.ErrorResponseItem{
				Message: fmt.Sprintf("field: %s, kind: %s, type: %s, should be equal: %s, actual value: %s", err.Field(), err.Kind(), err.Type(), err.Param(), err.Value()),
				Path:    err.Namespace(),
			})
		}

		return responseErrors
	}

	return nil
}
