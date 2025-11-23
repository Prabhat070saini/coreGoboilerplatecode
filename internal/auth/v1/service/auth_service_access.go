package authService

import (
	httpClient "github.com/example/testing/shared/lib/http"
	"github.com/example/testing/config"
	"github.com/example/testing/shared/clients/cache/cacheConfig"
)

type AuthServiceAccess struct {
	CacheService cacheConfig.Cache
	Config       *config.Env
	HttpService  *httpClient.HttpClientImpl
}

func NewAuthServiceAccess(cacheService cacheConfig.Cache, cfg *config.Env, httpService *httpClient.HttpClientImpl) *AuthServiceAccess {
	return &AuthServiceAccess{
		CacheService: cacheService,
		Config:       cfg,
		HttpService:  httpService,
	}
}
