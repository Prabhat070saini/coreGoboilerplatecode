package authHandler

import (
	"github.com/gin-gonic/gin"

	"github.com/example/testing/common/response"
	"github.com/example/testing/internal/auth/v1/dto"
	authService "github.com/example/testing/internal/auth/v1/service"
	middleware "github.com/example/testing/apis/middlewares"
)

type AuthHandlerMethods interface {
	Login(ctx *gin.Context)
}

type authHandler struct {
	authService authService.AuthServiceMethods
}

func NewAuthHandler(authService authService.AuthServiceMethods) *authHandler {
	return &authHandler{authService: authService}
}

func (ah *authHandler) Login(ctx *gin.Context) {

	// FIX: Correct way to get RequestContext
	reqCtx := middleware.GetReqContext(ctx)

	ip := ctx.ClientIP()
	userAgent := ctx.GetHeader("User-Agent")

	var req dto.LoginDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		exc := dto.GetLoginDtoValidationError(err)
		resp := response.ServiceOutput[struct{}]{ Exception: exc }
		response.SendRestResponse(ctx, &resp)
		return
	}

	output := ah.authService.Login(reqCtx.Ctx, &req, &ip, &userAgent)
	response.SendRestResponse(ctx, &output)
}
