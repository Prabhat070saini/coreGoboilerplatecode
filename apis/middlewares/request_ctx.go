package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

// RequestContext carries context info across layers
type RequestContext struct {
	Ctx context.Context
}



// Helper to get RequestContext from Gin context
func GetReqContext(c *gin.Context) *RequestContext {
	val, exists := c.Get("req_context")
	if !exists {
		// If not exists, create a new RequestContext
		reqCtx := &RequestContext{
			Ctx: c.Request.Context(),
		}
		c.Set("req_context", reqCtx)
		return reqCtx
	}
	return val.(*RequestContext)
}

