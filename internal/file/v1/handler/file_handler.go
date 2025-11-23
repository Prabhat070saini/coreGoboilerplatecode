package filehandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/example/testing/shared/response"
	"github.com/example/testing/internal/file/v1/dto"
	fileservice "github.com/example/testing/internal/file/v1/service/file_service"
	"github.com/gin-gonic/gin"
)

type FileHandlerMethods interface {
	UploadFile(ginCtx *gin.Context, ctx context.Context)
	FetchFile(ginCtx *gin.Context, ctx context.Context)
}

type fileHandler struct {
	fileService fileservice.FileServiceMethods
}

func NewFileHandler(fileService fileservice.FileServiceMethods) FileHandlerMethods {
	return &fileHandler{fileService: fileService}
}

func (f *fileHandler) UploadFile(ginCtx *gin.Context, ctx context.Context) {
	fmt.Println("Debug: UploadFile called")
	// reqCtx := middleware.GetReqContext(ctx)
	// logger.Info(reqCtx.Ctx, "file upload endpoint called")
	var req dto.UploadFileDto
	ginCtx.ShouldBind(&req) //nolint:errcheck

	// check for missing file
	if req.File == nil {
		output := &response.ServiceOutput[struct{}]{
			Exception: &response.Exception{
				Code:           500,
				Message:        "File is required",
				HttpStatusCode: http.StatusInternalServerError,
			},
		}
		response.SendRestResponse(ginCtx, output)
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
		response.SendRestResponse(ginCtx, output)
		return
	}

	// Check if folder == registration AND sessionId is missing in form fields
	if req.Folder == "registration" {
		if form, err := ginCtx.MultipartForm(); err == nil {
			if _, exists := form.Value["sessionId"]; !exists {
				output := &response.ServiceOutput[struct{}]{
					Exception: &response.Exception{
						Code:           500,
						Message:        "SessionId is required for registration folder",
						HttpStatusCode: http.StatusInternalServerError,
					},
				}
				response.SendRestResponse(ginCtx, output)
				return
			}
		}
	}

	// logger.Info(reqCtx.Ctx, "file upload request received", zap.String("folder", req.Folder), zap.String("filename", req.File.Filename), zap.String("session_id", req.SessionID))
	// Continue to service layer
	output := f.fileService.UploadFile(ctx, &req)
	response.SendRestResponse(ginCtx, &output)
}

func (f *fileHandler) FetchFile(ginCtx *gin.Context, ctx context.Context) {
	
	var req dto.FetchFileDto
	if err := ginCtx.ShouldBindQuery(&req); err != nil {
		exc := dto.GetFetchFileDtoValidationError(err)
		resp := response.ServiceOutput[struct{}]{
			Exception: exc,
		}
		response.SendRestResponse(ginCtx, &resp)
		return
	}

	output := f.fileService.FetchFile(ctx, &req)
	response.SendRestResponse(ginCtx, &output)
}
