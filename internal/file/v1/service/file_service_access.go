package fileServiceAccess


import (
	httpClient "github.com/example/testing/common/lib/http"
	"github.com/example/testing/config"
	"github.com/example/testing/shared/cache/cacheConfig"
	fileSystem "github.com/example/testing/shared/fileSystem"

)

type FileServiceAccess struct {
	CacheService cacheConfig.Cache
	Config       *config.Env
	HttpService  *httpClient.HttpClientImpl
	FileService  *fileSystem.FileService

}

func NewFileServiceAccess(cacheService cacheConfig.Cache, config *config.Env, httpService *httpClient.HttpClientImpl, fileService *fileSystem.FileService) *FileServiceAccess {
	return &FileServiceAccess{
		CacheService: cacheService,
		Config:       config,
		HttpService:  httpService,
		FileService:  fileService,
	}
}
