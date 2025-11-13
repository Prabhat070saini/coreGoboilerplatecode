package validator
import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9.]+$`)





func UsernameValidator(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	length := len(username)

	return length > 0 && length <= 106 && usernameRegex.MatchString(username)
}