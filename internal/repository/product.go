package repository

import (
	"example.com/internal/domain"
	"example.com/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// WithTx returns a new repository using the given transactional db.
func (r *ProductRepository) WithTx(tx *gorm.DB) domain.ProductRepository {
	return &ProductRepository{db: tx}
}

func (r *ProductRepository) Create(product models.Product) error {
	return r.db.Create(&product).Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	return products, r.db.Find(&products).Error
}

func (r *ProductRepository) GetByID(id uint) (models.Product, error) {
	var product models.Product
	return product, r.db.First(&product, id).Error
}

func (r *ProductRepository) Update(product models.Product) (models.Product, error) {
	return product, r.db.Save(&product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
