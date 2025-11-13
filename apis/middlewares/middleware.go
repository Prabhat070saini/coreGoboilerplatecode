package middleware

import (
	"github.com/example/testing/config"
	"github.com/example/testing/shared/cache/cacheConfig"
	"github.com/gin-gonic/gin"
	"github.com/example/testing/common/response"
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
	ctx.AbortWithStatusJSON(status, response.ApiResponse[any]{
		Code:    code,
		Message: message,
	})
}
