package authService

import "github.com/google/uuid"

type LoginResponse struct {
	Id                uuid.UUID `json:"id"`
	AccessToken       string    `json:"accessToken"`
	RefreshToken      string    `json:"refreshToken"`
	Name              string    `json:"name"`
}
