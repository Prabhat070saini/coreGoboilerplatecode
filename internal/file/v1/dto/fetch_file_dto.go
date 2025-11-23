package dto

import (
	"errors"

	"github.com/example/testing/shared/constants/exception"
	"github.com/example/testing/shared/response"
	"github.com/go-playground/validator/v10"
)

type FetchFileDto struct {
	Key string `form:"key" binding:"required"` // Changed from json to form
}

var fetchFileDtoErrorMap = map[string]map[string]*response.Exception{
	"Key": {
		"required": exception.GetException(exception.File_KEY_REQUIRED),
	},
}

func GetFetchFileDtoValidationError(err error) *response.Exception {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			if tagMap, ok := fetchFileDtoErrorMap[fe.Field()]; ok {
				if exc, ok := tagMap[fe.Tag()]; ok {
					return exc
				}
			}
		}
	}
	return exception.GetException(exception.INVALID_PAYLOAD)
}
