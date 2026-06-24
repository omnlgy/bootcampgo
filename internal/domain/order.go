package domain

import (
	"example.com/internal/models"
	"gorm.io/gorm"
)

type Order interface {
	CreateOrder() (models.Order, error)
}

type OrderRepository interface {
	Create(order models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id uint) (models.Order, error)
	Update(order models.Order) (models.Order, error)
	Delete(id uint) error
	WithTx(tx *gorm.DB) OrderRepository
}
