package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PGConfig struct {
	Host                     string
	Port                     int
	User                     string
	Password                 string
	DBName                   string
	SSLMode                  string
	MaxIdleConns             int
	MaxOpenConns             int
	ConnMaxLifetimeMinutes   int
	Logging                  bool
}

// NewDBConnection connects to PostgreSQL with best-practice settings.
func NewDBConnection(cfg *PGConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)

	// Configure logger
	logLevel := logger.Silent
	if cfg.Logging {
		logLevel = logger.Info
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("postgres connection failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.ConnMaxLifetimeMinutes))

	return db, nil
}

// PingDBConnection checks DB health
func PingDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	log.Println("✅ DB connection pinged successfully")
	return sqlDB.Ping()
}

// CloseDBConnection gracefully closes the DB
func CloseDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
