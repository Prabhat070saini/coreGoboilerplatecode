package filehandler

import (
	"fmt"
	"net/http"

	// middleware "github.com/example/testing/apis/middlewares"
	middleware "github.com/example/testing/apis/middlewares"
	"github.com/example/testing/common/lib/logger"
	"github.com/example/testing/common/response"
	"github.com/example/testing/internal/file/v1/dto"
	fileservice "github.com/example/testing/internal/file/v1/service/file_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileHandlerMethods interface {
	UploadFile(ctx *gin.Context)
	FetchFile(ctx *gin.Context)
}

type fileHandler struct {
	fileService fileservice.FileServiceMethods
}

func NewFileHandler(fileService fileservice.FileServiceMethods) FileHandlerMethods {
	return &fileHandler{fileService: fileService}
}

func (f *fileHandler) UploadFile(ctx *gin.Context) {
	fmt.Println("Debug: UploadFile called")
	reqCtx := middleware.GetReqContext(ctx)
	logger.Info(reqCtx.Ctx, "file upload endpoint called")
	var req dto.UploadFileDto
	ctx.ShouldBind(&req) //nolint:errcheck

	// check for missing file
	if req.File == nil {
		output := &response.ServiceOutput[struct{}]{
			Exception: &response.Exception{
				Code:           500,
				Message:        "File is required",
				HttpStatusCode: http.StatusInternalServerError,
			},
		}
		response.SendRestResponse(ctx, output)
		return
	}

	// check folder string
	if req.Folder == "" {
		output := &response.ServiceOutput[struct{}]{
			Exception: &response.Exception{
				Code:           500,
				Message:        "Folder is required",
				HttpStatusCode: http.StatusInternalServerError,
			},
		}
		response.SendRestResponse(ctx, output)
		return
	}

	// Check if folder == registration AND sessionId is missing in form fields
	if req.Folder == "registration" {
		if form, err := ctx.MultipartForm(); err == nil {
			if _, exists := form.Value["sessionId"]; !exists {
				output := &response.ServiceOutput[struct{}]{
					Exception: &response.Exception{
						Code:           500,
						Message:        "SessionId is required for registration folder",
						HttpStatusCode: http.StatusInternalServerError,
					},	
				}
				response.SendRestResponse(ctx, output)
				return
			}
		}
	}

	logger.Info(reqCtx.Ctx, "file upload request received", zap.String("folder", req.Folder), zap.String("filename", req.File.Filename), zap.String("session_id", req.SessionID))
	// Continue to service layer
	output := f.fileService.UploadFile(reqCtx.Ctx, &req)
	response.SendRestResponse(ctx, &output)
}

func (f *fileHandler) FetchFile(ctx *gin.Context) {
	reqCtx := middleware.GetReqContext(ctx)
	var req dto.FetchFileDto
	if err := ctx.ShouldBindQuery(&req); err != nil {
		exc := dto.GetFetchFileDtoValidationError(err)
		resp := response.ServiceOutput[struct{}]{
			Exception: exc,
		}
		response.SendRestResponse(ctx, &resp)
		return
	}

	output := f.fileService.FetchFile(reqCtx.Ctx, &req)
	response.SendRestResponse(ctx, &output)
}
