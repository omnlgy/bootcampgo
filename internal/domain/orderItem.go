package domain

import (
	"example.com/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(orderItem models.OrderItem) error
	GetByID(id uint) (models.OrderItem, error)
	Update(orderItem models.OrderItem) (models.OrderItem, error)
	Delete(id uint) error
	WithTx(tx *gorm.DB) OrderItemRepository
}
