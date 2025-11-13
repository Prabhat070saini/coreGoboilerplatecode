package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/example/testing/shared/fileSystem/fileConfig"
)

// s3Client implements FileProvider without logging
type s3Client struct {
	client   *s3.Client
	uploader *manager.Uploader
	config   fileConfig.S3Config
}

// NewS3Service creates a new AWS S3 service and validates credentials
func NewS3Service(cfg fileConfig.S3Config) (fileConfig.FileProvider, error) {
	// Validate required fields
	if cfg.AccessKeyID == "" || cfg.SecretAccessKey == "" {
		return nil, fmt.Errorf("❌ AWS credentials are required")
	}
	if cfg.Region == "" {
		return nil, fmt.Errorf("❌ AWS region is required")
	}
	if cfg.BucketName == "" {
		return nil, fmt.Errorf("❌ S3 bucket name is required")
	}

	awsCfg := aws.Config{
		Region:      cfg.Region,
		Credentials: credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
	}

	// Create S3 client
	client := s3.NewFromConfig(awsCfg)

	// Test credentials by making a simple API call
	_, err := client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to validate AWS credentials: %v", err)
	}

	// Check if the specified bucket exists and is accessible
	_, err = client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: aws.String(cfg.BucketName),
	})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to access S3 bucket '%s': %v", cfg.BucketName, err)
	}

	uploader := manager.NewUploader(client)

	return &s3Client{
		client:   client,
		uploader: uploader,
		config:   cfg,
	}, nil
}

// UploadFile uploads a file to S3 and returns the accessible URL as string
func (s *s3Client) UploadFile(ctx context.Context, file *multipart.FileHeader, folder *string, sessionId ...string) (*string, *string, error) {
	if file == nil {
		return nil, nil, fmt.Errorf("file not provided")
	}
	if file.Size > s.config.MaxFileSize {
		return nil, nil, fmt.Errorf("file size exceeds limit")
	}

	src, err := file.Open()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	key := s.generateKey(file.Filename, *folder, sessionId...)
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(fileBytes)
	}

	_, err = s.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.config.BucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPrivate,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to upload file: %w", err)
	}

	// Always return a presigned URL by default
	url, err := s.GenerateURL(ctx, &key, false)
	if err != nil {
		return nil, nil, fmt.Errorf("file uploaded but failed to generate URL: %w", err)
	}

	return url, &key, nil
}

// GenerateURL returns either a presigned URL (temporary) or permanent URL as *string
func (s *s3Client) GenerateURL(ctx context.Context, key *string, permanent bool) (*string, error) {
	if key == nil {
		return nil, fmt.Errorf("key not provided")
	}

	if permanent {
		// Permanent URL for public objects
		url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.config.BucketName, s.config.Region, *key)
		return &url, nil
	}

	// Generate presigned URL for private objects
	psClient := s3.NewPresignClient(s.client)
	req, err := psClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.config.BucketName),
		Key:    aws.String(*key),
	}, s3.WithPresignExpires(time.Duration(s.config.SignedURLExpiry)*time.Minute))
	if err != nil {
		return nil, fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return &req.URL, nil
}

// generateKey creates a unique S3 key with folder + sessionId + timestamp
func (s *s3Client) generateKey(filename string, folder string, sessionId ...string) string {
	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	name = strings.ReplaceAll(name, " ", "_")

	folderPath := folder
	if len(sessionId) > 0 && sessionId[0] != "" {
		folderPath = strings.TrimSuffix(folderPath, "/") + "/" + sessionId[0]
	}
	if folderPath != "" && !strings.HasSuffix(folderPath, "/") {
		folderPath += "/"
	}

	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s%s_%d%s", folderPath, name, timestamp, ext)
}
