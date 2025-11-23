package jwt

import (
	"errors"
	"log"
	"time"


	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken      = errors.New("invalid JWT token")
	ErrExpiredToken      = errors.New("token has expired")
	ErrInvalidSigning    = errors.New("unexpected signing method")
	ErrTokenTypeMismatch = errors.New("token type mismatch")
	ErrConfigNull        = errors.New("JWT configuration is not initialized")
)
type TokenType string

func GenerateJwtToken[T any](tokenType TokenType, payload T, expTimeInMinutes int, secret string) (*string, error) {
	claims := jwt.MapClaims{
		"payload": payload,
		"type":    tokenType,
		"exp":     time.Now().Add(time.Minute * time.Duration(expTimeInMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// ✅ Must capture error from SignedString
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil,err
	}

	return &signed, nil
}


func ValidateJwtToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	// secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Panic("JWT SECRET is not set in .env file")
	}

	// Parse token
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Ensure signing method is expected
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigning
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	// Extract and assert claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken

	}

	return claims, nil
}
