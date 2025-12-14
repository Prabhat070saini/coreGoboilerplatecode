package v1

import (
	middleware "github.com/example/testing/apis/middlewares"
	"github.com/example/testing/shared/constants"
	"github.com/example/testing/internal/initializer"
	"github.com/gin-gonic/gin"
)

const (
	FileGroupPrefix = "v1/file"
)

type FileRoutes struct{}

func NewFileRoutes(baseHandler *initializer.BaseHandler,router *gin.RouterGroup, middleware *middleware.Middlewares) *FileRoutes {
	protectedFile := router.Group(FileGroupPrefix)

	protectedFile.POST(
		"/upload",
		middleware.AuthMiddleware.AuthMiddleware(),
		middleware.PermissionMiddleware.PermissionMiddleware(constants.SuperDoctor, constants.Doctor),
		middleware.ContextInjectorMethods.InjectContext(baseHandler.FileHandler.UploadFile),
	)

	protectedFile.GET(
		"/fetch",
		middleware.ContextInjectorMethods.InjectContext(baseHandler.FileHandler.FetchFile),
	)

	return &FileRoutes{}
}
