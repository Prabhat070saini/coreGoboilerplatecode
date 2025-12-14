package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/example/testing/shared/constants"
	"github.com/example/testing/shared/constants/exception"
	"github.com/example/testing/shared/jwt"
	"github.com/example/testing/shared/lib/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type authMiddleware struct {
	access *MiddlewareAccess
}

type AuthMiddlewareMethods interface {
	AuthMiddleware() gin.HandlerFunc
}

func NewAuthMiddleware(access *MiddlewareAccess) AuthMiddlewareMethods {
	return &authMiddleware{
		access: access,
	}
}
func toStringSlice(val interface{}) []string {
	raw, ok := val.([]interface{})
	if !ok {
		return nil
	}

	result := make([]string, 0, len(raw))
	for _, v := range raw {
		if s, ok := v.(string); ok {
			result = append(result, strings.ToUpper(s)) // normalize
		}
	}
	return result
}

func (m *authMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			SendErrorResponse(c, http.StatusUnauthorized, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		token := authHeader[7:]

		tokenPayload, err := jwt.ValidateJwtToken(token, m.access.Cfg.JWT.AccessTokenSecret)
		if err != nil {
			logger.Error(c.Request.Context(), "ValidateJwtTokenFail", zap.Error(err))
			SendErrorResponse(c, exception.GetException(exception.INVALID_OR_EXPIRED_TOKEN).Code,
				exception.GetException(exception.INVALID_OR_EXPIRED_TOKEN).Message,
				http.StatusUnauthorized)
			return
		}

		payload, ok := tokenPayload["payload"].(map[string]interface{})
		if !ok {
			logger.Debug(c.Request.Context(), "fail to extract payload")
			SendErrorResponse(c, http.StatusUnauthorized, exception.GetException(exception.INTERNAL_SERVER_ERROR).Message, http.StatusInternalServerError)
			return
		}

		userId, ok := payload["id"].(string)
		if !ok || userId == "" {
			logger.Debug(c.Request.Context(), "User Id not found in token")
			SendErrorResponse(c, http.StatusUnauthorized, constants.UnauthorizedAccess, http.StatusUnauthorized)
			return
		}
		redisKey := fmt.Sprintf(constants.LoginAccessTokenRedisKey, userId)
		storeToken, err := m.access.cacheService.Get(c, redisKey)
		if err != nil || storeToken != token {
			logger.Debug(c.Request.Context(), "Store token not equal to access token", zap.Error(err))
			SendErrorResponse(c, constants.InvalidTokenLogout, exception.GetException(exception.INVALID_TOKEN_LOGOUT).Message, http.StatusUnauthorized)
			return
		}
		if roles, exists := payload["roles"]; exists {
			c.Set("roles", toStringSlice(roles)) // always []string
		}
		c.Set("userId", userId)

		c.Next()
	}
}
