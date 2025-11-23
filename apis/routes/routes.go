package routes

import (
	"time"

	"github.com/example/testing/shared/lib/logger"
	middleware "github.com/example/testing/apis/middlewares"
	v1 "github.com/example/testing/apis/routes/v1"
	"github.com/example/testing/config"
	"github.com/example/testing/internal/initializer"
	"github.com/gin-gonic/gin"
)

type Routes struct{}

func NewRoutes(router *gin.Engine, cfg *config.Env, baseHandler *initializer.BaseHandler, middlewares *middleware.Middlewares) {

	apiKeyCfg := &middleware.Config{
		APIKey: cfg.ApiKey,
	}
	apiKeyMiddleware := middleware.NewApiMiddleware(apiKeyCfg)

	corsCfg := &middleware.CorsConfig{
		Env:              cfg.AppEnv,
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowHeaders:     cfg.Cors.AllowHeaders,
		ExposeHeaders:    cfg.Cors.ExposeHeaders,
		AllowCredentials: cfg.Cors.AllowCredentials,
		MaxAge:           time.Duration(cfg.Cors.MaxAge) * time.Hour,
	}

	
	// Apply globally or per group
	router.Use(middleware.NewCorsMiddleware(corsCfg))

	appRoutes := router.Group("/api/your-service")
	appRoutes.Use(apiKeyMiddleware.Handler())

	appRoutes.GET("/ping", func(c *gin.Context) {
		logger.Debug(c.Request.Context(), "checking the value in the value")
		c.JSON(200, gin.H{"message": "pong"})
	})

	// -------------------------------
	// 5️⃣  Example for future modular routes
	// -------------------------------
	v1.NewAuthRoutes(baseHandler, appRoutes, middlewares)
	v1.NewFileRoutes(baseHandler, appRoutes, middlewares)
}
