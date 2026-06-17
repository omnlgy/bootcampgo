package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/omnlgy/bootcampgo/internal/domain"
	"github.com/omnlgy/bootcampgo/internal/logger"
	"github.com/omnlgy/bootcampgo/internal/model"
)

type CartService struct {
	Itmes map[string]int
}

func NewCartService() *CartService {
	return &CartService{
		Itmes: make(map[string]int),
	}
}

func (c *CartService) AddItem(item model.Product, quantity int) error {
	if err := item.Validate(); err != nil {
		return err
	}

	if item.Stock < quantity {
		return fmt.Errorf("stock tidak mencukupi")
	}

	c.Itmes[item.SKU] += quantity
	logger.LogTransaction(item.SKU, item.Price*float64(quantity))
	return nil
}

func (c *CartService) Checkout(products *map[string]model.Product, shipper domain.Shipper) (string, float64, error) {
	if len(c.Itmes) == 0 {
		return "", 0, fmt.Errorf("keranjang kosong")
	}

	var totalAmount float64

	for sku, quantity := range c.Itmes {
		product, exists := (*products)[sku]
		if !exists {
			return "", 0, fmt.Errorf("produk dengan SKU %s tidak ditemukan", sku)
		}
		totalAmount += product.Price * float64(quantity)
		product.Stock -= quantity
		(*products)[sku] = product
	}

	orderId, err := uuid.NewV7()
	if err != nil {
		return "", 0, err
	}

	c.Itmes = make(map[string]int)

	fmt.Printf("%s Cost: Rp %.2f\n", shipper.GetCourierName(), shipper.CalculateCost())

	return orderId.String(), totalAmount, nil
}
