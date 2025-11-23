package dto

import (
	"errors"

	"github.com/example/testing/shared/constants/exception"
	"github.com/example/testing/shared/response"
	"github.com/go-playground/validator/v10"
)

type LoginDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

var loginValidationErrorMap = map[string]map[string]*response.Exception{
	"Email": {
		"email":    exception.GetException(exception.INVALID_EMAIL),
		"required": exception.GetException(exception.EMAIL_REQUIRED),
	},
	"Password": {
		"required": exception.GetException(exception.PASSWORD_REQUIRED),
	},
}

func GetLoginDtoValidationError(err error) *response.Exception {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			if tagMap, ok := loginValidationErrorMap[fe.Field()]; ok {
				if exc, ok := tagMap[fe.Tag()]; ok {
					return exc
				}
			}
		}
	}
	return exception.GetException(exception.INVALID_PAYLOAD)
}
