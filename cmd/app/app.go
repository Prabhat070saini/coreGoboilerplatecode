package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	middleware "github.com/example/testing/apis/middlewares"
	"github.com/example/testing/apis/routes"
	"github.com/example/testing/shared/constants"
	httpClient "github.com/example/testing/shared/lib/http"
	"github.com/example/testing/shared/lib/logger"
	"github.com/example/testing/shared/validator"
	"github.com/example/testing/config"
	"github.com/example/testing/internal/initializer"
	"github.com/example/testing/shared/clients/cache"
	"github.com/example/testing/shared/clients/cache/cacheConfig"
	"github.com/example/testing/shared/clients/database"
	fileSystem "github.com/example/testing/shared/clients/fileSystem"
	"github.com/example/testing/shared/clients/fileSystem/fileConfig"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	cfg          *config.Env
	router       *gin.Engine
	server       *http.Server
	db           *gorm.DB
	cacheService cacheConfig.Cache
	repo         *initializer.BaseRepository
	fileService  *fileSystem.FileService
	service      *initializer.BaseService
	handler      *initializer.BaseHandler
	httpService  *httpClient.HttpClientImpl
	middleware   *middleware.Middlewares
}

func NewApp(cfg *config.Env) *App {
	app := &App{
		cfg: cfg,
	}
	app.initialize()
	return app
}

func (a *App) initialize() {

	// Initialize logger
	logger.Init(logger.LogConfig{
		Level:            a.cfg.Log.Level,
		Format:           a.cfg.Log.Format,
		EnableCaller:     a.cfg.Log.EnableCaller,
		EnableStacktrace: a.cfg.Log.EnableStacktrace,
		RequestIDKey:     constants.RequestIDKey,
	})

	// Connect to the database 
	dbConnection, err := database.NewDBConnection(&database.PGConfig{
		Host:                   a.cfg.DB.Host,
		Port:                   a.cfg.DB.Port,
		User:                   a.cfg.DB.User,
		Password:               a.cfg.DB.Password,
		DBName:                 a.cfg.DB.DBName,
		MaxIdleConns:           a.cfg.DB.MaxIdleConnection,
		SSLMode:                a.cfg.DB.SSLMode,
		MaxOpenConns:           a.cfg.DB.MaxOpenConnection,
		ConnMaxLifetimeMinutes: a.cfg.DB.ConnectionLifeTimeMinute,
		Logging:                a.cfg.DB.Logging,
	})

	if err != nil {
		logger.Error(context.Background(), "database connection failed")
		panic(err)

	}
	if err := database.PingDBConnection(dbConnection); err != nil {
		logger.Error(context.Background(), "DB unhealthy:", zap.Error(err))
	}

	a.db = dbConnection
	// Connect Redis
	cacheConfiguration := &cacheConfig.Config{
		Driver:   a.cfg.Cache.Driver,
		Addr:     a.cfg.Cache.Addr,
		Password: a.cfg.Cache.Password,
		DB:       a.cfg.Cache.Db,
	}
	cacheConnection, err := cache.Init(cacheConfiguration)
	if err != nil {
		panic(err)

	}
	a.cacheService = cacheConnection

	httpService := httpClient.Init(httpClient.HttpConfig{
		Timeout:      10 * time.Second,
		RequestIDKey: constants.RequestIDKey,
	})
	a.httpService = httpService

	s3Cfg := fileConfig.S3Config{
		Region:          a.cfg.S3.Region,
		AccessKeyID:     a.cfg.S3.AccessKeyID,
		SecretAccessKey: a.cfg.S3.SecretAccessKey,
		BucketName:      a.cfg.S3.BucketName,
		SignedURLExpiry: a.cfg.S3.SignedURLExpiry,
		MaxFileSize:     int64(a.cfg.S3.MaxFileSize), // 5MB
	}

	_, fileErr := fileSystem.Initialize(fileSystem.S3, s3Cfg)
	if fileErr != nil {
		panic(fileErr)
	}

	fs := fileSystem.GetInstance()
	a.fileService = fs

	// Init Router
	a.router = gin.Default()
	a.middleware = middleware.NewMiddlewares(a.cfg, a.cacheService)

	a.router.Use(a.middleware.TracingMiddleware.TracingMiddleware())
	a.router.Use(a.middleware.SecurityMiddleware.SecurityHeadersMiddleware(a.cfg.AppEnv))

	a.repo = initializer.NewBaseRepository(a.db, a.cacheService, a.cfg)
	a.service = initializer.NewBaseService(a.cacheService, a.cfg, a.repo, a.db, a.httpService,a.fileService)
	a.handler = initializer.NewBaseHandler(a.cfg, a.service)

	routes.NewRoutes(a.router, a.cfg, a.handler, a.middleware)
	validator.RegisterValidations()
	a.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", a.cfg.Port),
		Handler:      a.router,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
	}

	a.registerRoutes()
}
func (a *App) registerRoutes() {
	// Example route
	a.router.GET("/health", func(c *gin.Context) {
		logger.Debug(c.Request.Context(), "health")
		// a.cacheService.Set(context.Background(), "greeting:valuechecking", "Hello Prabhat", 10*time.Second)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

func (a *App) Run() {
	// Start Gin HTTP server
	go func() {
		logger.Info(context.Background(), "starting HTTP server", zap.String("addr", a.server.Addr))
		fmt.Printf("Starting application... http://localhost:%d",a.cfg.Port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(context.Background(), "HTTP server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Signal(syscall.SIGTERM))
	<-quit
	logger.Info(context.Background(), "Received shutdown signal, shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := a.server.Shutdown(ctx); err != nil {
		logger.Error(context.Background(), "HTTP server forced to shutdown", zap.Error(err))
	} else {
		logger.Info(context.Background(), "HTTP server stopped gracefully")
	}

	// Shutdown DB
	if err := database.CloseDBConnection(a.db); err != nil {
		logger.Error(context.Background(), "Database close error", zap.Error(err))
	} else {
		logger.Info(context.Background(), "Database connection closed gracefully")
	}

	// Shutdown Redis
	if err := cache.Close(); err != nil {
		logger.Error(context.Background(), "Redis close error", zap.Error(err))
	} else {
		logger.Info(context.Background(), "Redis connection closed gracefully")
	}
}
