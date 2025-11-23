package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
)


type HandlerFuncWithCtx func(c *gin.Context, stdCtx context.Context)


type requestCtxMiddleware struct {
}

type ContextInjectorMethods interface {
	InjectContext(handler HandlerFuncWithCtx) gin.HandlerFunc
}

func NewRequestCtxMiddleware() ContextInjectorMethods {
	return &requestCtxMiddleware{}
}


func (m *requestCtxMiddleware) InjectContext(handler HandlerFuncWithCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get existing context or create and set it
		reqCtx, exists := c.Get("req_context")
		var ctx context.Context

		if !exists {
			ctx = c.Request.Context()
			c.Set("req_context", ctx)
		} else {
			ctx = reqCtx.(context.Context)
		}

		handler(c, ctx)
	}
}
