package file

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"sync"

	"github.com/example/testing/shared/clients/fileSystem/fileConfig"
	"github.com/example/testing/shared/clients/fileSystem/s3"
)

type ProviderType string

const (
	S3 ProviderType = "s3"
	// Future: GCP, Azure, local FS, etc.
)

type FileService struct {
	provider fileConfig.FileProvider
}

var (
	instance *FileService
	once     sync.Once
)

// Initialize singleton with chosen provider
func Initialize(provider ProviderType, cfg any) (*FileService, error) {
	var err error
	once.Do(func() {
		switch provider {
		case S3:
			s3Cfg, ok := cfg.(fileConfig.S3Config)
			if !ok {
				err = fmt.Errorf("❌ invalid config type for S3 provider")
				log.Println(err)
				return
			}
			p, err2 := s3.NewS3Service(s3Cfg)
			if err2 != nil {
				err = fmt.Errorf("❌ failed to initialize S3 service: %v", err2)
				log.Println(err)
				return
			}
			instance = &FileService{provider: p}
			log.Println("✅ File service (S3) initialized successfully")
		default:
			err = fmt.Errorf("❌ unsupported file provider: %s", provider)
			log.Println(err)
		}
	})

	if err != nil {
		log.Printf("❌ File service initialization failed: %v\n", err)
		return nil, err
	}

	return instance, nil
}

// GetInstance returns singleton
func GetInstance() *FileService {
	if instance == nil {
		// FIXME: uncomment below line if you want to throw error or want service not to be initialized without file system
		// log.Fatal("[FileService] Service not initialized. Call Initialize first.")
		log.Println("[FileService] Service not initialized. Call Initialize first.")
		return nil
	}
	return instance
}

// UploadFile delegates to provider
// UploadFile delegates to provider
func (f *FileService) UploadFile(ctx context.Context, file interface{}, folder string, sessionId ...string) (string, string, error) {
	fh, ok := file.(*multipart.FileHeader)
	if !ok {
		return "", "", fmt.Errorf("invalid file type")
	}

	urlPtr, keyPtr, err := f.provider.UploadFile(ctx, fh, &folder, sessionId...)
	if err != nil {
		return "", "", err
	}
	if urlPtr == nil {
		return "", "", fmt.Errorf("provider returned nil URL")
	}
	return *urlPtr, *keyPtr, nil
}

// GeneratePresignedURL delegates to provider
func (f *FileService) GenerateURL(ctx context.Context, key string, permanent bool) (string, error) {
	urlPtr, err := f.provider.GenerateURL(ctx, &key, permanent)
	if err != nil {
		return "", err
	}
	if urlPtr == nil {
		return "", fmt.Errorf("provider returned nil URL")
	}
	return *urlPtr, nil
}
