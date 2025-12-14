package v1

import (
	middleware "github.com/example/testing/apis/middlewares"
	"github.com/example/testing/shared/lib/logger"
	"github.com/example/testing/internal/initializer"
	"github.com/gin-gonic/gin"
)

const (
	authBasePath = "/v1/auth"
)

type AuthRoutes struct {
}

func NewAuthRoutes(baseHandler *initializer.BaseHandler, router *gin.RouterGroup, middleware *middleware.Middlewares) *AuthRoutes {
	protectedAuth := router.Group(authBasePath)
	protectedAuth.POST("/login", middleware.ContextInjectorMethods.InjectContext(baseHandler.AuthHandler.Login))
	protectedAuth.POST("/forgot-password", func(c *gin.Context) {
		logger.Debug(c.Request.Context(), "checking the value in the value")
		c.JSON(200, gin.H{"message": "pong"})
	})

	return &AuthRoutes{}
}
