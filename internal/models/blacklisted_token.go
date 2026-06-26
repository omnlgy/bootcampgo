package models

import "time"

type BlacklistedToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"type:text;not null;uniqueIndex" json:"-"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (BlacklistedToken) TableName() string {
	return "blacklisted_tokens"
}
