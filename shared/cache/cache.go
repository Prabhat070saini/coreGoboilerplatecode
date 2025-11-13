package cache

import (
	"fmt"
	"sync"

	"github.com/example/testing/shared/cache/cacheConfig"
	"github.com/example/testing/shared/cache/redis"
)

var (
	instance cacheConfig.Cache
	once     sync.Once
)

func Init(cfg *cacheConfig.Config) (cacheConfig.Cache, error) {
	var err error
	once.Do(func() {
		switch cfg.Driver {
		case "redis":
			instance, err = redis.NewRedisCache(cfg)

		case "memcache":
			// instance, err = memcache.NewMemcacheCache(cfg)
		case "memory":
			// instance, err = memory.NewMemoryCache(cfg)
		default:
			err = fmt.Errorf("unsupported cache driver: %s", cfg.Driver)
		}
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Your", cfg.Driver, "Cache connected successfully")
	return instance, nil
}

func GetCache() (cacheConfig.Cache, error) {
	if instance == nil {
		return nil, fmt.Errorf("cache not initialized")
	}
	return instance, nil
}

func Close() error {
	if instance != nil {
		return instance.Close()
	}
	return nil
}
