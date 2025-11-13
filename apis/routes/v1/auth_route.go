package v1

import (
	middleware "github.com/example/testing/apis/middlewares"
	"github.com/example/testing/common/lib/logger"
	"github.com/example/testing/internal/initializer"
	"github.com/gin-gonic/gin"
)

const (
	authBasePath = "/v1/auth"
)

type AuthRoutes struct {
}

func NewAuthRoutes(baseHandler *initializer.BaseHandler, routerGroup *gin.RouterGroup, middleware *middleware.Middlewares) *AuthRoutes {
	authGroup := routerGroup.Group(authBasePath)
	authGroup.POST("/login", baseHandler.AuthHandler.Login)
	authGroup.POST("/forgot-password", func(c *gin.Context) {
		logger.Debug(c.Request.Context(), "checking the value in the value")
		c.JSON(200, gin.H{"message": "pong"})
	})

	return &AuthRoutes{}
}
