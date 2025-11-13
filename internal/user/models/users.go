package models


import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID           int            `gorm:"primaryKey;column:id"`
	UUID         uuid.UUID      `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();not null"`
	Name         *string        `gorm:"column:name;type:varchar(100)"`
	Email        string         `gorm:"column:email;type:varchar(255);not null;uniqueIndex"`
	PasswordHash string         `gorm:"column:password_hash;type:text;not null"`
	IsActive     bool           `gorm:"column:is_active;not null;default:true"`
	IsBlocked    bool           `gorm:"column:is_blocked;not null;default:false"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:current_timestamp"`
	CreatedBy    *int           `gorm:"column:created_by;default:-1"`
	ModifiedAt   *time.Time     `gorm:"column:modified_at"`
	ModifiedBy   *int           `gorm:"column:modified_by"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy    *int           `gorm:"column:deleted_by"`
}

// TableName sets the actual PostgreSQL table name with schema
func (Users) TableName() string {
	return "users"
}
