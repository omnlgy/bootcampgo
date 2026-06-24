package repository

import (
	"example.com/internal/domain"
	"example.com/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// WithTx returns a new repository using the given transactional db.
func (r *OrderRepository) WithTx(tx *gorm.DB) domain.OrderRepository {
	return &OrderRepository{db: tx}
}

func (r *OrderRepository) Create(order models.Order) error {
	return r.db.Create(&order).Error
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	return orders, r.db.Find(&orders).Error
}

func (r *OrderRepository) GetByID(id uint) (models.Order, error) {
	var order models.Order
	return order, r.db.First(&order, id).Error
}

func (r *OrderRepository) Update(order models.Order) (models.Order, error) {
	return order, r.db.Save(&order).Error
}

func (r *OrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
