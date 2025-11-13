package constants

import (
	"errors"

	"github.com/example/testing/common/jwt"
)



const (
	AccessToken  jwt.TokenType = "access_token"
	RefreshToken jwt.TokenType = "refresh_token"
)

var (
	ErrInvalidToken   = errors.New("invalid token")
	ErrExpiredToken   = errors.New("token has expired")
	ErrInvalidSigning = errors.New("invalid signing method")
)
