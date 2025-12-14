package ratelimiter

import (
	"context"
	"strconv"
	"time"

	"github.com/example/testing/shared/clients/cache/cacheConfig"
	"github.com/gin-gonic/gin"
)

type RateLimiter interface {
	Allow(ctx context.Context, key string) (bool, error)
	KeyFunc(c *gin.Context) string
}





type RedisClientLimiter struct {
	cache      cacheConfig.Cache
	refillRate float64
	burst      float64
	keyFunc    func(*gin.Context) string
}

func NewRedisClientLimiter(cache cacheConfig.Cache, refillRate float64, burst float64, keyFunc func(*gin.Context) string) *RedisClientLimiter {
	return &RedisClientLimiter{
		cache:      cache,
		refillRate: refillRate,
		burst:      burst,
		keyFunc:    keyFunc,
	}
}

func (l *RedisClientLimiter) KeyFunc(c *gin.Context) string {
	return l.keyFunc(c)
}

func (l *RedisClientLimiter) Allow(ctx context.Context, key string) (bool, error) {

	now := float64(time.Now().Unix())

	tokenKey := key + ":tokens"
	tsKey := key + ":ts"

	tokenStr, _ := l.cache.Get(ctx, tokenKey)
	tsStr, _ := l.cache.Get(ctx, tsKey)

	var tokens float64
	var lastTs float64

	if tokenStr == "" || tsStr == "" {
		tokens = l.burst
		lastTs = now
	} else {
		tokens, _ = strconv.ParseFloat(tokenStr, 64)
		lastTs, _ = strconv.ParseFloat(tsStr, 64)
	}

	// refill logic
	elapsed := now - lastTs
	refilled := elapsed * l.refillRate
	tokens = min(tokens+refilled, l.burst)
	lastTs = now

	if tokens < 1 {
		_ = l.cache.SetWithExp(ctx, tokenKey, strconv.FormatFloat(tokens, 'f', 4, 64), time.Minute)
		_ = l.cache.SetWithExp(ctx, tsKey, strconv.FormatFloat(lastTs, 'f', 4, 64), time.Minute)
		return false, nil
	}

	// consume token
	tokens -= 1

	_ = l.cache.SetWithExp(ctx, tokenKey, strconv.FormatFloat(tokens, 'f', 4, 64), time.Minute)
	_ = l.cache.SetWithExp(ctx, tsKey, strconv.FormatFloat(lastTs, 'f', 4, 64), time.Minute)

	return true, nil
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
