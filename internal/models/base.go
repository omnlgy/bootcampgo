package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base provides the common fields for all models: UUID primary key
// with DB-side uuid_v7() default, plus GORM-managed timestamps and
// soft-delete support.
type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_v7()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
