package authHandler

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/example/testing/shared/response"
	"github.com/example/testing/internal/auth/v1/dto"
	authService "github.com/example/testing/internal/auth/v1/service"
)

type AuthHandlerMethods interface {
	Login(ginCtx *gin.Context, ctx context.Context)
}

type authHandler struct {
	authService authService.AuthServiceMethods
}

func NewAuthHandler(authService authService.AuthServiceMethods) *authHandler {
	return &authHandler{authService: authService}
}

func (ah *authHandler) Login(ginCtx *gin.Context, ctx context.Context) {

	ip := ginCtx.ClientIP()
	userAgent := ginCtx.GetHeader("User-Agent")

	var req dto.LoginDto
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		exc := dto.GetLoginDtoValidationError(err)
		resp := response.ServiceOutput[struct{}]{Exception: exc}
		response.SendRestResponse(ginCtx, &resp)
		return
	}

	output := ah.authService.Login(ctx, &req, &ip, &userAgent)
	response.SendRestResponse(ginCtx, &output)
}
