package userRepository

import (
	"github.com/example/testing/config"
	"github.com/example/testing/shared/cache/cacheConfig"
	"gorm.io/gorm"
)

type UserRepositoryAccess struct {
	DB           *gorm.DB
	cacheService cacheConfig.Cache
	Config       *config.Env
}

func NewUserRepoAccess(db *gorm.DB, cacheService cacheConfig.Cache, cfg *config.Env) *UserRepositoryAccess {
	return &UserRepositoryAccess{
		DB:           db,
		cacheService: cacheService,
		Config:       cfg,
	}
}
