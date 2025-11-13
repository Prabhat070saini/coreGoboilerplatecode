package fileservice

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"io"
// 	"net/http"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"

// 	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-user-service/internal/file/v1/dto"
// 	fileservice "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-user-service/internal/file/v1/service"
// 	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-user-service/shared/constants"
// 	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-user-service/shared/constants/exception"
// 	"go.uber.org/zap"
// )

// // --- Mock S3 Client ---
// type MockS3Client struct {
// 	mock.Mock
// }

// func (m *MockS3Client) UploadFile(ctx context.Context, file io.Reader, folder string, sessionID string) (*fileservice.UploadFileResponse, error) {
// 	args := m.Called(ctx, file, folder, sessionID)
// 	return args.Get(0).(*fileservice.UploadFileResponse), args.Error(1)
// }

// func (m *MockS3Client) GeneratePresignedURL(ctx context.Context, key string) (*fileservice.FetchFileResponse, error) {
// 	args := m.Called(ctx, key)
// 	return args.Get(0).(*fileservice.FetchFileResponse), args.Error(1)
// }

// // --- Mock Logger ---
// type MockLogger struct{}

// func (m *MockLogger) Error(msg string, fields ...zap.Field) {}
// func (m *MockLogger) Warn(msg string, fields ...zap.Field)  {}
// func (m *MockLogger) Info(msg string, fields ...zap.Field)  {}
// func (m *MockLogger) Debug(msg string, fields ...zap.Field) {}

// // --- Test Suite ---
// type FileServiceTestSuite struct {
// 	suite.Suite
// 	service fileservice.FileServiceMethods
// 	mockS3  *MockS3Client
// }

// func (suite *FileServiceTestSuite) SetupTest() {
// 	suite.mockS3 = new(MockS3Client)
// 	access := &fileservice.FileServiceAccess{
// 		S3:     suite.mockS3,
// 		Logger: &MockLogger{},
// 	}
// 	suite.service = fileservice.NewFileService(access)
// }

// // --- UploadFile Tests ---
// func (suite *FileServiceTestSuite) TestUploadFile_Success() {
// 	ctx := context.Background()
// 	file := bytes.NewBufferString("dummy file content")
// 	dto := &dto.UploadFileDto{File: file, Folder: "docs", SessionID: "sess123"}
// 	resp := &fileservice.UploadFileResponse{Key: "docs/file.pdf", URL: "https://s3.example.com/docs/file.pdf"}

// 	suite.mockS3.On("UploadFile", ctx, file, "docs", "sess123").Return(resp, nil)

// 	out := suite.service.UploadFile(ctx, dto)
// 	assert.Equal(suite.T(), http.StatusOK, out.HttpStatusCode)
// 	assert.Equal(suite.T(), constants.UploadFileSuccess, out.Message)
// 	assert.Equal(suite.T(), resp.Key, out.OutputData.Key)
// }

// func (suite *FileServiceTestSuite) TestUploadFile_FileTooLarge() {
// 	ctx := context.Background()
// 	file := bytes.NewBufferString("too large content")
// 	dto := &dto.UploadFileDto{File: file, Folder: "img", SessionID: "sess999"}

// 	suite.mockS3.On("UploadFile", ctx, file, "img", "sess999").Return(nil, errors.New("file size exceeds limit"))

// 	out := suite.service.UploadFile(ctx, dto)
// 	assert.Equal(suite.T(), http.StatusRequestEntityTooLarge, out.HttpStatusCode)
// 	assert.Contains(suite.T(), out.Message, "file size exceeds")
// }

// // --- FetchFile Tests ---
// func (suite *FileServiceTestSuite) TestFetchFile_Success() {
// 	ctx := context.Background()
// 	key := "docs/file.pdf"
// 	resp := &fileservice.FetchFileResponse{URL: "https://s3.example.com/docs/file.pdf"}

// 	suite.mockS3.On("GeneratePresignedURL", ctx, key).Return(resp, nil)

// 	out := suite.service.FetchFile(ctx, &dto.FetchFileDto{Key: key})
// 	assert.Equal(suite.T(), http.StatusOK, out.HttpStatusCode)
// 	assert.Equal(suite.T(), constants.FileFetched, out.Message)
// 	assert.Equal(suite.T(), resp.URL, out.OutputData.URL)
// }

// func (suite *FileServiceTestSuite) TestFetchFile_Error() {
// 	ctx := context.Background()
// 	key := "invalid/key"
// 	suite.mockS3.On("GeneratePresignedURL", ctx, key).Return(nil, errors.New("key not found"))

// 	out := suite.service.FetchFile(ctx, &dto.FetchFileDto{Key: key})
// 	assert.Equal(suite.T(), http.StatusInternalServerError, out.HttpStatusCode)
// 	assert.Equal(suite.T(), exception.INTERNAL_SERVER_ERROR.Message, out.Message)
// }

// func TestFileServiceTestSuite(t *testing.T) {
// 	suite.Run(t, new(FileServiceTestSuite))
// }
