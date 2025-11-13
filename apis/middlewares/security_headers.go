package middleware

import "github.com/gin-gonic/gin"

type securityHeaderMiddleware struct {
	access *MiddlewareAccess
}

type SecurityHeadersMiddlewareMethods interface {
	SecurityHeadersMiddleware(env string) gin.HandlerFunc
}

func NewSecurityHeaderMiddleware(access *MiddlewareAccess) SecurityHeadersMiddlewareMethods {
	return &securityHeaderMiddleware{
		access: access,
	}
}

// SecurityHeadersMiddleware sets secure HTTP headers for all responses.
func (m *securityHeaderMiddleware) SecurityHeadersMiddleware(env string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// HTTPS-only headers
		if c.Request.TLS != nil {
			c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		}

		// General security headers
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
		c.Writer.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'")

		c.Next()
	}
}
