package middleware

import (
	"github.com/example/testing/config"
	"github.com/example/testing/shared/clients/cache/cacheConfig"
	"github.com/example/testing/shared/response"
	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	AuthMiddleware         AuthMiddlewareMethods
	SecurityMiddleware     SecurityHeadersMiddlewareMethods
	TracingMiddleware      TracingMiddlewareMethods
	PermissionMiddleware   PermissionMiddlewareMethods
	ContextInjectorMethods ContextInjectorMethods
	ApiKeyMiddleware       ApiKeyMiddlewareMethods
}

type MiddlewareAccess struct {
	Cfg          *config.Env
	cacheService cacheConfig.Cache
}

func NewMiddlewares(cfg *config.Env, cacheService cacheConfig.Cache) *Middlewares {
	access := &MiddlewareAccess{
		Cfg: cfg, cacheService: cacheService,
	}
	apiKeyCfg := &Config{
		APIKey: cfg.ApiKey,
	}
	return &Middlewares{
		AuthMiddleware:         NewAuthMiddleware(access),
		SecurityMiddleware:     NewSecurityHeaderMiddleware(access),
		TracingMiddleware:      NewTracingMiddleware(access),
		PermissionMiddleware:   NewPermissionMiddleware(access),
		ContextInjectorMethods: NewRequestCtxMiddleware(),
		ApiKeyMiddleware:       NewApiMiddleware(apiKeyCfg),
	}
}

func SendErrorResponse(ctx *gin.Context, code int, message string, status int) {
	ctx.AbortWithStatusJSON(status, response.ApiResponse[any]{
		Code:    code,
		Message: message,
	})
}
