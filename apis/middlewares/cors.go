package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config contains all customizable CORS options
type CorsConfig struct {
	Env              string
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           time.Duration
}

// New returns a configured CORS middleware handler
func NewCorsMiddleware(cfg *CorsConfig) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     cfg.AllowMethods,
		AllowHeaders:     cfg.AllowHeaders,
		ExposeHeaders:    cfg.ExposeHeaders,
		AllowCredentials: cfg.AllowCredentials,
		MaxAge:           cfg.MaxAge,
	})
}
