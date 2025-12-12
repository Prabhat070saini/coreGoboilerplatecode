package middleware

import (
	"net/http"

	"github.com/example/testing/shared/constants"
	"github.com/gin-gonic/gin"
)

type ApiKeyMiddlewareMethods interface {
	Handler() gin.HandlerFunc
	Public() gin.HandlerFunc
}

// Config holds the API key for the middleware
type Config struct {
	APIKey string
}

// Middleware struct
type ApiKeyMiddleware struct {
	cfg *Config
}

// New creates a new instance of the API key middleware
func NewApiMiddleware(cfg *Config) *ApiKeyMiddleware {
	return &ApiKeyMiddleware{
		cfg: cfg,
	}
}

// Handler returns the Gin middleware handler
func (m *ApiKeyMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if this route is marked as public
		if skip, ok := c.Get(constants.SkipAPIKeyCheck); ok && skip.(bool) {
			c.Next()
			return
		}

		apiKeyHeader := c.GetHeader("x-api-key")

		if apiKeyHeader == "" || apiKeyHeader != m.cfg.APIKey {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Invalid API key",
				},
			)
			return
		}

		c.Next()
	}
}

func (m *ApiKeyMiddleware) Public() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.SkipAPIKeyCheck, true)
		c.Next()
	}
}
