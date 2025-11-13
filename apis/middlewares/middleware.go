package middleware

import (
	"github.com/example/testing/config"
	"github.com/example/testing/shared/cache/cacheConfig"
	"github.com/example/testing/common/constants"
	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	AuthMiddleware     AuthMiddlewareMethods
	SecurityMiddleware SecurityHeadersMiddlewareMethods
	TracingMiddleware  TracingMiddlewareMethods
}

type MiddlewareAccess struct {
	Cfg          *config.Env
	cacheService cacheConfig.Cache
}

func NewMiddlewares(cfg *config.Env, cacheService cacheConfig.Cache) *Middlewares {
	access := &MiddlewareAccess{
		Cfg: cfg, cacheService: cacheService,
	}

	return &Middlewares{
		AuthMiddleware:     NewAuthMiddleware(access),
		SecurityMiddleware: NewSecurityHeaderMiddleware(access),
		TracingMiddleware:  NewTracingMiddleware(access),
	}
}

func SendErrorResponse(ctx *gin.Context, code int, message string, status int) {
	ctx.AbortWithStatusJSON(status, constants.ApiResponse[any]{
		Code:    code,
		Message: message,
	})
}
