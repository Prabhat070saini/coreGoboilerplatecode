package middleware


import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		apiKeyHeader := c.GetHeader("x-api-key")

		if apiKeyHeader == "" || apiKeyHeader != m.cfg.APIKey {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid API key",
				},
			)
			return
		}

		c.Next()
	}
}
