package repository

import (
	"example.com/internal/domain"
	"example.com/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Create(orderItem models.OrderItem) error {
	return r.db.Create(&orderItem).Error
}

func (r *OrderItemRepository) GetByID(id uint) (models.OrderItem, error) {
	var orderItem models.OrderItem
	err := r.db.First(&orderItem, id).Error
	return orderItem, err
}

func (r *OrderItemRepository) Update(orderItem models.OrderItem) (models.OrderItem, error) {
	err := r.db.Save(&orderItem).Error
	return orderItem, err
}

func (r *OrderItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}

func (r *OrderItemRepository) WithTx(tx *gorm.DB) domain.OrderItemRepository {
	return &OrderItemRepository{db: tx}
}
