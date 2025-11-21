package initializer

import (
	"github.com/example/testing/config"
	authHandler "github.com/example/testing/internal/auth/v1/handler"
	fileHandler "github.com/example/testing/internal/file/v1/handler"
)

type BaseHandler struct {
	// UserHandler userHandler.UserHandlerMethods
	
	FileHandler fileHandler.FileHandlerMethods
	AuthHandler authHandler.AuthHandlerMethods
}

func NewBaseHandler(cfg *config.Env, baseService *BaseService) *BaseHandler {
	return &BaseHandler{
		// UserHandler: userHandler.NewUserHandler(baseService.UserService),
		FileHandler: fileHandler.NewFileHandler(baseService.FileService),
		AuthHandler: authHandler.NewAuthHandler(baseService.AuthService),
	}
}
