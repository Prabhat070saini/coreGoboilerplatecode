package logger

import (
	"context"
	
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logInstance      *zap.Logger
	once             sync.Once
	// globalRequestCtx context.Context
	enabledLevels    zapcore.LevelEnabler
)

type LogConfig struct {
	Level            string `mapstructure:"level" validate:"required"`
	Format           string `mapstructure:"format" validate:"required"`
	EnableCaller     bool   `mapstructure:"enable_caller"`
	EnableStacktrace bool   `mapstructure:"enable_stacktrace"`
	RequestIDKey     interface{}
}

var contextKeyForRequestID interface{}

// ---- Initialization ----

func Init(cfg LogConfig) *zap.Logger {
	once.Do(func() {
		logInstance = buildLogger(cfg)
		enabledLevels = getLevelEnabler(cfg.Level)
		contextKeyForRequestID = cfg.RequestIDKey
	})
	return logInstance
}

func buildLogger(cfg LogConfig) *zap.Logger {
	level := parseLogLevel(cfg.Level)
	encoder := getEncoder(cfg.Format)

	core := zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		level,
	)

	var options []zap.Option
	if cfg.EnableCaller {
		options = append(options, zap.AddCaller())
	}
	if cfg.EnableStacktrace {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	return zap.New(core, options...)
}

func parseLogLevel(level string) zapcore.Level {
	var lvl zapcore.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		return zapcore.InfoLevel
	}
	return lvl
}

func getEncoder(format string) zapcore.Encoder {
	if strings.ToLower(format) == "json" {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		return zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderCfg)
}

func getLevelEnabler(level string) zapcore.LevelEnabler {
	lvl := parseLogLevel(level)
	return zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= lvl
	})
}

// ---- Context Management ----

// func SetRequestContext(ctx context.Context) {
// 	globalRequestCtx = ctx
// }

// func GetRequestContext() context.Context {
// 	return globalRequestCtx
// }

// ---- Logger Access ----

func Get() *zap.Logger {
	if logInstance == nil {
		return Init(LogConfig{
			Level:            "info",
			Format:           "console",
			EnableCaller:     true,
			EnableStacktrace: false,
		})
	}
	return logInstance
}

func FromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Get()
	}

	logger := Get()
	
	if contextKeyForRequestID != nil {
		if requestID := ctx.Value(contextKeyForRequestID); requestID != nil {
			logger = logger.With(zap.String("tracing_id", requestID.(string)))
		}
	}
	return logger
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	if enabled(enabledLevels, zapcore.InfoLevel) {
		FromContext(ctx).Info(msg, fields...)
	}
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	if enabled(enabledLevels, zapcore.ErrorLevel) {
		FromContext(ctx).Error(msg, fields...)
	}
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if enabled(enabledLevels, zapcore.WarnLevel) {
		FromContext(ctx).Warn(msg, fields...)
	}
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if enabled(enabledLevels, zapcore.DebugLevel) {
		FromContext(ctx).Debug(msg, fields...)
	}
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if enabled(enabledLevels, zapcore.FatalLevel) {
		FromContext(ctx).Fatal(msg, fields...)
	}
}

func enabled(enabler zapcore.LevelEnabler, level zapcore.Level) bool {
	if enabler == nil {
		return true
	}
	return enabler.Enabled(level)
}

// ---- Context Logger ----

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		ctx = context.TODO()
	}
	if contextKeyForRequestID != nil {
		if requestID := ctx.Value(contextKeyForRequestID); requestID != nil {
			return Get().With(zap.String("tracing_id", requestID.(string)))
		}
	}
	return Get()
}
