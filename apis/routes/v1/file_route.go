package v1

import (
	middleware "github.com/example/testing/apis/middlewares"
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
		baseHandler.FileHandler.UploadFile,
	)

	routerGroup.GET(
		FileGroupPrefix+"/fetch",
		// middleware.AuthMiddlewareMethods.AuthMiddleware(),
		baseHandler.FileHandler.FetchFile,
	)

	return &FileRoutes{}
}
