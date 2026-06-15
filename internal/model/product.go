package model

import "fmt"

type Product struct {
	SKU   string
	Name  string
	Price float64
	Stock int
}

func (p *Product) Validate() error {
	if p.SKU == "" {
		return fmt.Errorf("sku is required")
	}
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if p.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return nil
}
