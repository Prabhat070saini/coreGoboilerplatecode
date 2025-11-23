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

func NewFileRoutes(baseHandler *initializer.BaseHandler, routerGroup *gin.RouterGroup, middleware *middleware.Middlewares) *FileRoutes {

	routerGroup.POST(
		FileGroupPrefix+"/upload",
		middleware.AuthMiddleware.AuthMiddleware(),
		middleware.PermissionMiddleware.PermissionMiddleware(constants.SuperDoctor, constants.Doctor),
		middleware.ContextInjectorMethods.InjectContext(baseHandler.FileHandler.UploadFile),
	)

	routerGroup.GET(
		FileGroupPrefix+"/fetch",
		middleware.ContextInjectorMethods.InjectContext(baseHandler.FileHandler.FetchFile),
	)

	return &FileRoutes{}
}
