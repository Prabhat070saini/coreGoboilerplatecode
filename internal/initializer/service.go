package initializer

import (
	"github.com/example/testing/shared/cache/cacheConfig"
	file "github.com/example/testing/shared/fileSystem"

	httpClient "github.com/example/testing/common/lib/http"
	"github.com/example/testing/config"
	authService "github.com/example/testing/internal/auth/v1/service"
	fileServiceAccess "github.com/example/testing/internal/file/v1/service"
	fileservice "github.com/example/testing/internal/file/v1/service/file_service"
	userService "github.com/example/testing/internal/user/v1/service"
	"gorm.io/gorm"
)

type BaseService struct {
	UserService userService.UserServiceMethods
	AuthService authService.AuthServiceMethods
	FileService fileservice.FileServiceMethods
}

func NewBaseService(cacheService cacheConfig.Cache, cfg *config.Env, baseRepo *BaseRepository, db *gorm.DB, httpService *httpClient.HttpClientImpl,fileService *file.FileService) *BaseService {
	userServiceAccess := userService.NewUserServiceAccess(cacheService, cfg, db)
	fileServiceAccess := fileServiceAccess.NewFileServiceAccess(cacheService, cfg, httpService, fileService)
	// Services
	authServiceAccess := authService.NewAuthServiceAccess(cacheService, cfg, httpService)
	userSvc := userService.NewUserService(baseRepo.UserRepository, userServiceAccess)
	fileSvc := fileservice.NewFileService(fileServiceAccess)
	return &BaseService{
		UserService: userSvc,
		AuthService: authService.NewAuthService(authServiceAccess, userSvc),
		FileService: fileSvc,
	}
}
