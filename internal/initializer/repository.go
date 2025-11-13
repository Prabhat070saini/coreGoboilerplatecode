package initializer

import (
	"github.com/example/testing/config"
	userRepository "github.com/example/testing/internal/user/repository"
	"github.com/example/testing/shared/cache/cacheConfig"
	"gorm.io/gorm"
)

type BaseRepository struct {
	UserRepository userRepository.UserRepositoryMethods
	// AuditLog  auditrepository.AuditLogRepositoryMethods
}

// NewBaseRepository wires all individual repositories using shared access
func NewBaseRepository(db *gorm.DB, cacheService cacheConfig.Cache, cfg *config.Env) *BaseRepository {
	// authRepoAccess := authrepositoryaccess.NewAuthRepoAccess(db, cacheService, cfg)
	userRepoAccess := userRepository.NewUserRepoAccess(db, cacheService, cfg)

	return &BaseRepository{
		UserRepository: userRepository.NewUserRepository(userRepoAccess),
		// UserLogin: userloginrepository.NewUserLoginRepository(authRepoAccess),
	}
}
