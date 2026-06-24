package domain

import (
	"example.com/internal/models"
	"gorm.io/gorm"
)

type Product interface{}

type ProductRepository interface {
	Create(product models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id uint) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id uint) error
	WithTx(tx *gorm.DB) ProductRepository
}
