package domain

import (
	"example.com/internal/models"
	"gorm.io/gorm"
)

type BlacklistedTokenRepository interface {
	Create(token models.BlacklistedToken) error
	Exists(tokenHash string) (bool, error)
	WithTx(tx *gorm.DB) BlacklistedTokenRepository
}
