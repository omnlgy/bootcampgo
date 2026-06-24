package service

import (
	"example.com/internal/domain"
	"example.com/internal/models"
)

type ProductService struct {
	productRepo domain.ProductRepository
}

func NewProductService(productRepo domain.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) CreateProduct(product models.Product) error {
	return s.productRepo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *ProductService) UpdateProduct(product models.Product) (models.Product, error) {
	return s.productRepo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}
