package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/example/testing/shared/cache/cacheConfig"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(cfg *cacheConfig.Config) (cacheConfig.Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &redisCache{client: rdb}, nil
}

// Implement Cache interface
func (r *redisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
func (r *redisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
func (r *redisCache) Exists(ctx context.Context, key string) (bool, error) {
	val, err := r.client.Exists(ctx, key).Result()
	return val > 0, err
}
func (r *redisCache) Close() error {
	err := r.client.Close()
	if err != nil {
		return err
	}
	fmt.Println("Redis cache closed successfully")
	return nil
}
