package middleware

import (
	"context"

	"github.com/example/testing/shared/lib/logger"
	"github.com/example/testing/shared/constants"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type tracingMiddleware struct {
	access *MiddlewareAccess
}

type TracingMiddlewareMethods interface {
	TracingMiddleware() gin.HandlerFunc
}

func NewTracingMiddleware(access *MiddlewareAccess) TracingMiddlewareMethods {
	return &tracingMiddleware{
		access: access,
	}
}
func (m *tracingMiddleware) TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer for request duration
		startTime := time.Now()

		// Get or generate request ID
		requestID := getOrGenerateRequestID(c)

		// Create context with request ID and start time
		ctx := enrichContext(c, requestID, startTime)

		// Set up logger with request context
		requestLogger := setupRequestLogger(ctx)

		// Set response headers
		setResponseHeaders(c, requestID)

		// Store values in Gin context
		storeInGinContext(c, requestID, requestLogger, startTime)

		// Log request start
		logRequestStart(requestLogger, c, startTime)

		// Process request
		c.Next()

		// Log request completion
		logRequestCompletion(requestLogger, c, startTime)
	}
}

// Helper functions

func getOrGenerateRequestID(c *gin.Context) string {
	requestID := c.GetHeader("tracing_id")
	if strings.TrimSpace(requestID) == "" {
		requestID = uuid.New().String()
	}
	return requestID
}
func enrichContext(c *gin.Context, requestID string, startTime time.Time) context.Context {
	ctx := c.Request.Context()

	// Add request ID to context
	ctx = context.WithValue(ctx, constants.RequestIDKey, requestID)

	// Add start time to context
	ctx = context.WithValue(ctx, constants.RequestStartTimeKey, startTime)

	// Update the request with new context
	c.Request = c.Request.WithContext(ctx)
	return ctx
}

func setupRequestLogger(ctx context.Context) *zap.Logger {
	// Create logger with request context
	loggerWithCtx := logger.WithContext(ctx)

	// Add additional common fields if needed
	return loggerWithCtx
}

func setResponseHeaders(c *gin.Context, requestID string) {
	// Set request ID in response header

	c.Writer.Header().Set("tracing_id", requestID)
}

func storeInGinContext(c *gin.Context, requestID string, logger *zap.Logger, startTime time.Time) {
	// Store request ID in Gin context
	c.Set(string(constants.RequestIDKey), requestID)
	// Store logger in Gin context
	c.Set("logger", logger)

	// Store start time in Gin context
	c.Set("start_time", startTime)
}

func logRequestStart(logger *zap.Logger, c *gin.Context, startTime time.Time) {
	logger.Info("Request started",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.String("ip", c.ClientIP()),
		zap.String("user_agent", c.Request.UserAgent()),
		zap.Time("start_time", startTime),
	)
}

func logRequestCompletion(logger *zap.Logger, c *gin.Context, startTime time.Time) {
	duration := time.Since(startTime)
	status := c.Writer.Status()

	logger.Info("Request completed",
		zap.Int("status", status),
		zap.Duration("duration_ms", duration),
		zap.String("duration_human", duration.String()),
		zap.Time("end_time", time.Now()),
	)
}
