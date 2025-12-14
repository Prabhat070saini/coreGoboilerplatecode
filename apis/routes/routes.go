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
	corsCfg := &middleware.CorsConfig{
		Env:              cfg.AppEnv,
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowHeaders:     cfg.Cors.AllowHeaders,
		ExposeHeaders:    cfg.Cors.ExposeHeaders,
		AllowCredentials: cfg.Cors.AllowCredentials,
		MaxAge:           time.Duration(cfg.Cors.MaxAge) * time.Hour,
	}

	
	// Apply globally 
	router.Use(middleware.NewCorsMiddleware(corsCfg))


	root := router.Group("/api/auth-service")
	// not need api key
	openRouter := root.Group("/")
	// need api key
	privateRouter := root.Group("/")
	privateRouter.Use(middlewares.ApiKeyMiddleware.Handler())

	openRouter.GET("/ping", func(c *gin.Context) {
		logger.Debug(c.Request.Context(), "checking the value in the value")
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1.NewAuthRoutes(baseHandler, privateRouter, middlewares)
	v1.NewFileRoutes(baseHandler, privateRouter, middlewares)
}
