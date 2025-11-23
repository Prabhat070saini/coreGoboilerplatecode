package validator

// PasswordStrengthValidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)
	digitRegex       = regexp.MustCompile(`[0-9]`)
	specialCharRegex = regexp.MustCompile(`[\W_]`)
)

func PasswordStrengthValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	length := len(password)

	return length >= 8 && length <= 16 &&
		uppercaseRegex.MatchString(password) &&
		lowercaseRegex.MatchString(password) &&
		digitRegex.MatchString(password) &&
		specialCharRegex.MatchString(password)
}
