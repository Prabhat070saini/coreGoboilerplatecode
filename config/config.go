package config

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// DBConfig
type DBConfig struct {
	Driver                   string `mapstructure:"driver" validate:"required"`
	Host                     string `mapstructure:"host" validate:"required"`
	Port                     int    `mapstructure:"port" validate:"required"`
	User                     string `mapstructure:"user" validate:"required"`
	Password                 string `mapstructure:"password" validate:"required"`
	DBName                   string `mapstructure:"db_name" validate:"required"`
	SSLMode                  string `mapstructure:"ssl_mode" validate:"required"`
	MaxIdleConnection        int    `mapstructure:"max_idle_connection" validate:"required"`
	MaxOpenConnection        int    `mapstructure:"max_open_connection" validate:"required"`
	ConnectionLifeTimeMinute int    `mapstructure:"conn_max_lifetime_minutes" validate:"required"`
	Logging                  bool   `mapstructure:"logging"`
}

// JWTConfig
type JWTConfig struct {
	AccessTokenExpiryMin  int    `mapstructure:"access_token_expiry_min" validate:"required"`
	RefreshTokenExpiryMin int    `mapstructure:"refresh_token_expiry_min" validate:"required"`
	AccessTokenSecret     string `mapstructure:"access_token_secret" validate:"required"`
	RefreshTokenSecret    string `mapstructure:"refresh_token_secret" validate:"required"`
}

// LogConfig
type LogConfig struct {
	Level            string `mapstructure:"level" validate:"required"`
	Format           string `mapstructure:"format" validate:"required"`
	EnableCaller     bool   `mapstructure:"enable_caller"`
	EnableStacktrace bool   `mapstructure:"enable_stacktrace"`
}

// RedisConfig
type CacheConfig struct {
	Addr     string `mapstructure:"addr" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Db       int    `mapstructure:"db"`
	Driver   string `mapstructure:"driver" validate:"required"`
}

// Aws s3 config
type S3Config struct {
	Region          string `mapstructure:"aws_region" validate:"required"`
	AccessKeyID     string `mapstructure:"aws_access_key_id" validate:"required"`
	SecretAccessKey string `mapstructure:"aws_secret_access_key" validate:"required"`
	BucketName      string `mapstructure:"aws_s3_bucket_name" validate:"required"`
	SignedURLExpiry int    `mapstructure:"aws_signed_url_expiry" validate:"required"` // in minutes
	MaxFileSize     int    `mapstructure:"max_file_size" validate:"required"`         // In bytes
}
type CorsConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	ExposeHeaders    []string `mapstructure:"expose_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age_hours"`
}
type EmailConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Username string `mapstructure:"username" `
	Password string `mapstructure:"password" validate:"required"`
	From     string `mapstructure:"from" validate:"required"`
	Provider string `mapstructure:"provider" validate:"required"`
}

type ServiceBaseUrlConfig struct {
	CommunicationService string `mapstructure:"communication_service" validate:"required"`
}

type PasswordResetConfig struct {
	ResetLink    string `mapstructure:"reset_link" validate:"required"`
	ResetLinkExp int    `mapstructure:"reset_link_expiry_min" validate:"required"`
}
type CryptoConfig struct {
	AesKey string `mapstructure:"aes_key" validate:"required"`
}

// Env represents the full config
type Env struct {
	ApiKey         string               `mapstructure:"ApiKey" validate:"required"`
	AppName        string               `mapstructure:"AppName" validate:"required"`
	AppVersion     string               `mapstructure:"AppVersion" validate:"required"`
	BaseUrl        string               `mapstructure:"BaseUrl" validate:"required,url"`
	AppEnv         string               `mapstructure:"Environment" validate:"required"`
	Port           int                  `mapstructure:"ServerPort" validate:"required"`
	DB             DBConfig             `mapstructure:"db" validate:"required"`
	JWT            JWTConfig            `mapstructure:"jwt" validate:"required"`
	Log            LogConfig            `mapstructure:"log" validate:"required"`
	Cache          CacheConfig          `mapstructure:"cache" validate:"required"`
	S3             S3Config             `mapstructure:"aws" validate:"required"`
	ServiceBaseUrl ServiceBaseUrlConfig `mapstructure:"service_base_url" validate:"required"`
	PasswordReset  PasswordResetConfig  `mapstructure:"password_reset" validate:"required"`
	Crypto         CryptoConfig         `mapstructure:"crypto" validate:"required"`
	Cors           CorsConfig           `mapstructure:"cors" validate:"required"`
	Email          EmailConfig          `mapstructure:"email" validate:"required"`
}

var (
	cfg     *Env
	once    sync.Once
	val     = validator.New()
	loadErr error
)

func LoadConfig() (*Env, error) {
	once.Do(func() {
		env := Env{}

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")

		if err := viper.ReadInConfig(); err != nil {
			loadErr = fmt.Errorf("could not read config file: %w", err)
			return
		}

		if err := viper.Unmarshal(&env); err != nil {
			loadErr = fmt.Errorf("could not unmarshal config: %w", err)
			return
		}

		if err := val.Struct(env); err != nil {
			loadErr = fmt.Errorf("config validation failed: %w", err)
			return
		}

		cfg = &env
	})

	return cfg, loadErr
}
