package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordStrengthValidator", PasswordStrengthValidator) //nolint:errcheck
		v.RegisterValidation("usernameValidator", UsernameValidator)                 //nolint:errcheck
	}
}
