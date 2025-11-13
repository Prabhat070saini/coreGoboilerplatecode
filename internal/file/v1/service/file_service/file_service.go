package fileservice

// import (
// 	"context"
// 	"net/http"
// 	"strings"

// 	"github.com/example/testing/common/lib/logger"
// 	"github.com/example/testing/common/response"

// 	"github.com/example/testing/common/constants/exception"
// 	"github.com/example/testing/common/utils"
// 	"github.com/example/testing/internal/file/v1/dto"
// 	fileserviceaccess "github.com/example/testing/internal/file/v1/service"
// 	"go.uber.org/zap"
// )

// type FileServiceMethods interface {
// 	UploadFile(ctx context.Context, payload *dto.UploadFileDto) response.ServiceOutput[*UploadFileResponse]
// 	FetchFile(ctx context.Context, payload *dto.FetchFileDto) response.ServiceOutput[*FetchFileResponse]
// }

// type fileService struct {
// 	access *fileserviceaccess.FileServiceAccess
// }

// func NewFileService(access *fileserviceaccess.FileServiceAccess) FileServiceMethods {
// 	return &fileService{
// 		access: access,
// 	}
// }

// func (f *fileService) UploadFile(ctx context.Context, payload *dto.UploadFileDto) response.ServiceOutput[*UploadFileResponse] {
// 	// Step 4: Upload to S3
// 	url, key, err := f.access.FileService.UploadFile(ctx, payload.File, payload.Folder, payload.SessionID)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "file size exceeds") {
// 			ex := exception.GetException(exception.FILE_TOO_LARGE)
// 			ex.Message = err.Error()
// 			return response.ServiceOutput[*UploadFileResponse]{
// 				Exception: &response.Exception{
// 					Code:           ex.Code,
// 					Message:        ex.Message,
// 					HttpStatusCode: ex.HttpStatusCode,
// 				},
// 			}
// 		}
// 		// Generic error
// 		return response.ServiceOutput[*UploadFileResponse]{
// 			Exception: &response.Exception{
// 				Code:           500,
// 				Message:        err.Error(),
// 				HttpStatusCode: http.StatusInternalServerError,
// 			},
// 		}
// 	}

// 	// Step 5: Return success
// 	resp := &UploadFileResponse{
// 		Key: key,
// 		URL: url,
// 	}
// 	return response.ServiceOutput[*UploadFileResponse]{
// 		Success: &response.Success[*UploadFileResponse]{
// 			Code:           http.StatusOK,
// 			Message:        "File uploaded successfully",
// 			HttpStatusCode: http.StatusOK,
// 			Data:           resp,
// 		},
// 	}
// }

// func (f *fileService) FetchFile(ctx context.Context, payload *dto.FetchFileDto) response.ServiceOutput[*FetchFileResponse] {
// 	result, err := f.access.FileService.GenerateURL(ctx, payload.Key, false)
// 	if err != nil {
// 		logger.Error(ctx, "fail to fetch file", zap.Error(err))
// 		return utils.ServiceError[*FetchFileResponse](exception.INTERNAL_SERVER_ERROR)
// 	}

// 	logger.Info(ctx, "fetch file successfully", zap.Error(err))

// 	resp := &FetchFileResponse{URL: result}

// 	return response.ServiceOutput[*FetchFileResponse]{
// 		Success: &response.Success[*FetchFileResponse]{
// 			Code:           http.StatusOK,
// 			Message:        "File fetched successfully",
// 			HttpStatusCode: http.StatusOK,
// 			Data:           resp,
// 		},
// 	}
// }
