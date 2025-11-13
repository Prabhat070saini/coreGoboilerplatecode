package initializer

import (
	"github.com/example/testing/config"
	authHandler "github.com/example/testing/internal/auth/v1/handler"
)

type BaseHandler struct {
	// UserHandler userHandler.UserHandlerMethods
	AuthHandler authHandler.AuthHandlerMethods
}

func NewBaseHandler(cfg *config.Env, baseService *BaseService) *BaseHandler {
	return &BaseHandler{
		// UserHandler: userHandler.NewUserHandler(baseService.UserService),
		AuthHandler: authHandler.NewAuthHandler(baseService.AuthService),
	}
}
