package userService

import (
	"github.com/example/testing/shared/clients/cache/cacheConfig"
	"github.com/example/testing/config"
	"gorm.io/gorm"
)

type UserServiceAccess struct {
	CacheService  cacheConfig.Cache
	Config        *config.Env
	TransactionDb *gorm.DB
}

func NewUserServiceAccess(cacheService cacheConfig.Cache, cfg *config.Env, transactionDb *gorm.DB) *UserServiceAccess {
	return &UserServiceAccess{
		CacheService:  cacheService,
		Config:        cfg,
		TransactionDb: transactionDb,
	}
}
