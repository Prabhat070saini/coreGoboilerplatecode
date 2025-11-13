package fileConfig

import (
	"context"
	"mime/multipart"
)

// S3Config holds AWS S3 configuration
type S3Config struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	SignedURLExpiry int   // in minutes
	MaxFileSize     int64 // in bytes
}

// FileProvider defines the interface for uploading files and generating URLs
type FileProvider interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, folder *string, sessionId ...string) (*string,*string, error)
	GenerateURL(ctx context.Context, key *string, permanent bool) (*string, error) // added `permanent` flag
}
