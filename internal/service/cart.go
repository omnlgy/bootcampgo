package service

import (
	"fmt"

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
	return nil
}
