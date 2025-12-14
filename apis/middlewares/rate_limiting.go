package middleware

import (
	"context"
	"net/http"
	"github.com/example/testing/shared/ratelimiter"
	"github.com/gin-gonic/gin"
)

type RateLimitingMiddlewareMethods interface {
	RateLimit(limiter ratelimiter.RateLimiter) gin.HandlerFunc
}

type RateLimitingMiddleware struct{}

func NewRateLimitingMiddleware() RateLimitingMiddlewareMethods {
	return &RateLimitingMiddleware{}
}

func (m *RateLimitingMiddleware) RateLimit(limiter ratelimiter.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {

		key := limiter.KeyFunc(c)
		if key == "" {
			panic("rate limiting key is empty")
		}

		allowed, err := limiter.Allow(context.Background(), key)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "RateLimiter Error"})
			return
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		c.Next()
	}
}



/*

root := router.Group("/api/auth-service")

ipLimiter := ratelimiter.NewRedisClientLimiter(
	a.cacheService,
	10, // refillRate tokens per second
	20, // burst
	func(c *gin.Context) string {
		return c.ClientIP()
	},
)

root.Use(a.middleware.RateLimitingMiddleware.RateLimit(ipLimiter))











protectedLimiter := ratelimiter.NewRedisClientLimiter(
	a.cacheService,
	5,
	10,
	func(c *gin.Context) string {
		return c.GetString("userId")
	},
)

protected.Use(a.middleware.RateLimitingMiddleware.RateLimit(protectedLimiter))
















public.GET("/login",
    a.middleware.RateLimitingMiddleware.RateLimit(
        ratelimiter.NewRedisClientLimiter(
            a.cacheService,
            1,  // 1 token/sec
            3,  // burst 3
            func(c *gin.Context) string { 
                return "login:" + c.ClientIP()
            },
        ),
    ),
    handler.Login,
)

























globalLimiter := ratelimiter.NewRedisClientLimiter(
	a.cacheService,
	20,
	40,
	func(c *gin.Context) string { return c.ClientIP() },
)

a.router.Use(a.middleware.RateLimitingMiddleware.RateLimit(globalLimiter))











pkill -f tmp/app
pkill -f air



*/