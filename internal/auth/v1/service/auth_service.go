package authService

import (
	"context"

	"fmt"
	"time"

	"github.com/example/testing/common/hashing"
	"github.com/example/testing/common/jwt"
	"github.com/example/testing/internal/auth/v1/dto"

	"github.com/example/testing/common/constants"
	"github.com/example/testing/common/constants/exception"
	"github.com/example/testing/common/lib/logger"
	"github.com/example/testing/common/response"
	"github.com/example/testing/common/utils"
	userService "github.com/example/testing/internal/user/v1/service"
	"go.uber.org/zap"
)

type AuthServiceMethods interface {
	Login(ctx context.Context, payload *dto.LoginDto, ip *string, userAgent *string) response.ServiceOutput[*LoginResponse]
}

type authService struct {
	access      *AuthServiceAccess
	userService userService.UserServiceMethods
}

func NewAuthService(access *AuthServiceAccess, userService userService.UserServiceMethods) AuthServiceMethods {
	return &authService{
		access:      access,
		userService: userService,
	}
}

func (s *authService) Login(ctx context.Context, payload *dto.LoginDto, ip *string, userAgent *string) response.ServiceOutput[*LoginResponse] {

	selectFields := []string{"uuid", "password_hash", "is_blocked", "email", "name"}
	conditions := map[string]interface{}{"email": payload.Email}
	getUserOutput := s.userService.GetUser(ctx, conditions, selectFields...)
	// Map specific errors first …
	if getUserOutput.Exception != nil {
		return utils.HandleException[*LoginResponse](*getUserOutput.Exception)
	}

	if getUserOutput.Data == nil {
		logger.Info(ctx, "User Not Found", zap.String("not foun", "vlaue"))
		return utils.ServiceError[*LoginResponse](exception.INVALID_CREDENTIALS)
	}

	if getUserOutput.Data.IsBlocked {
		return utils.ServiceError[*LoginResponse](exception.USER_BLOCKED)
	}

	if !hashing.CompareHashAndPassword(&getUserOutput.Data.PasswordHash, &payload.Password) {
		logger.Debug(ctx,"password not	 match")
		return utils.ServiceError[*LoginResponse](exception.INVALID_CREDENTIALS)
	}

	// Prepare token payload (minimal, safe data only)
	accessTokenPayload := struct {
		Id string `json:"id"`
	}{
		Id: getUserOutput.Data.UUID.String(),
	}
	refreshTokenPayload := struct {
		Id string `json:"id"`
	}{
		Id: getUserOutput.Data.UUID.String(),
	}

	// Generate token

	accessToken, err := jwt.GenerateJwtToken(constants.AccessToken, accessTokenPayload, s.access.Config.JWT.AccessTokenExpiryMin, s.access.Config.JWT.AccessTokenSecret)
	if err != nil {
		logger.Error(ctx, "Failed to generate access Token")
		return utils.ServiceError[*LoginResponse](exception.INTERNAL_SERVER_ERROR)
	}

	refreshToken, err := jwt.GenerateJwtToken(constants.RefreshToken, refreshTokenPayload, s.access.Config.JWT.RefreshTokenExpiryMin, s.access.Config.JWT.RefreshTokenSecret)
	if err != nil {
		logger.Error(ctx, "Failed to generate refresh Token")
		return utils.ServiceError[*LoginResponse](exception.INTERNAL_SERVER_ERROR)
	}
	// Save token to Redis  LoginRefreshTokenRedisKey
	accessRedisKey := fmt.Sprintf(constants.LoginAccessTokenRedisKey, getUserOutput.Data.UUID)

	refreshRedisKey := fmt.Sprintf(constants.LoginRefreshTokenRedisKey, getUserOutput.Data.UUID)

	redisErr := s.access.CacheService.Set(ctx, accessRedisKey, accessToken, time.Minute*time.Duration(s.access.Config.JWT.AccessTokenExpiryMin))
	if redisErr != nil {
		return utils.ServiceError[*LoginResponse](exception.INTERNAL_SERVER_ERROR)
	}
	redisErr = s.access.CacheService.Set(ctx, refreshRedisKey, refreshToken, time.Minute*time.Duration(s.access.Config.JWT.RefreshTokenExpiryMin))
	if redisErr != nil {
		return utils.ServiceError[*LoginResponse](exception.INTERNAL_SERVER_ERROR)
	}

	resp := &LoginResponse{AccessToken: *accessToken, Id: getUserOutput.Data.UUID, RefreshToken: *refreshToken, Name: *getUserOutput.Data.Name}
	tokenPayload, err := jwt.ValidateJwtToken(*accessToken, "s.access.Config.JWT.AccessTokenSecret")
	if err != nil {
		logger.Error(ctx, "Failed to generate refresh Token", zap.Error(err))
		return utils.ServiceError[*LoginResponse](exception.INTERNAL_SERVER_ERROR)
	}
	fmt.Println(tokenPayload, "payload")






	return response.ServiceOutput[*LoginResponse]{
		Success: &response.Success[*LoginResponse]{
			Code:           200,
			Message:        constants.LoginSuccess,
			HttpStatusCode: 200,
			Data:           resp,
		},
	}
}
