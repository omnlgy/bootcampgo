package repository

import (
	"example.com/internal/domain"
	"example.com/internal/models"
	"gorm.io/gorm"
)

var _ domain.BlacklistedTokenRepository = (*BlacklistedTokenRepository)(nil)

type BlacklistedTokenRepository struct {
	db *gorm.DB
}

func NewBlacklistedTokenRepository(db *gorm.DB) *BlacklistedTokenRepository {
	return &BlacklistedTokenRepository{db: db}
}

func (r *BlacklistedTokenRepository) WithTx(tx *gorm.DB) domain.BlacklistedTokenRepository {
	return &BlacklistedTokenRepository{db: tx}
}

func (r *BlacklistedTokenRepository) Create(token models.BlacklistedToken) error {
	return r.db.Create(&token).Error
}

func (r *BlacklistedTokenRepository) Exists(tokenStr string) (bool, error) {
	var count int64
	err := r.db.Model(&models.BlacklistedToken{}).
		Where("token = ? AND expires_at > NOW()", tokenStr).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
