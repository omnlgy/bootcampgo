package main

import (
	"errors"
	"fmt"
)

type Address struct {
	Street string
	City   string
	Zip    string
}

type Vendor struct {
	Id      string
	Name    string
	Address Address
}

type Product struct {
	Sku    string
	Name   string
	Price  float64
	Stock  int
	Vendor *Vendor
}

func NewProduct(sku, name string, price float64, stock int, vendor *Vendor) (*Product, error) {
	if sku == "" {
		return nil, errors.New("sku cannot be empty")
	}
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}
	if stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}
	if vendor == nil {
		return nil, errors.New("vendor cannot be nil")
	}

	return &Product{
		Sku:    sku,
		Name:   name,
		Price:  price,
		Stock:  stock,
		Vendor: vendor,
	}, nil
}

func (p *Product) Restock(quantity int) error {
	if quantity < 0 {
		return errors.New("quantity must be positive")
	}
	p.Stock += quantity
	return nil
}

func (p *Product) DeductStock(quantity int) error {
	if quantity < 0 {
		return errors.New("quantity must be positive")
	}
	if p.Stock < quantity {
		return errors.New("not enough stock")
	}
	p.Stock -= quantity
	return nil
}

func (p Product) PrintDetails() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Product Details:")
	fmt.Printf("SKU         : %s\n", p.Sku)
	fmt.Printf("Name        : %s\n", p.Name)
	fmt.Printf("Price       : %.2f\n", p.Price)
	fmt.Printf("Stock       : %d\n", p.Stock)
	fmt.Printf("Dikirim Oleh: %s (%s, %s)\n", p.Vendor.Name, p.Vendor.Address.Street, p.Vendor.Address.City)
	fmt.Println("-------------------------------------------------------")
}

type ShippingAddress struct {
	Street string
	City   string
	Zip    string
}

type CustomerProfile struct {
	Name        string
	LoyaltyTier string
	TotalSpent  float64
	Address     *ShippingAddress
}

func (c *CustomerProfile) New(name string, address *ShippingAddress) {
	c.Name = name
	c.LoyaltyTier = "silver"
	c.TotalSpent = 0
	c.Address = address
}

func (c *CustomerProfile) RecordTransaction(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	c.TotalSpent += amount

	switch {
	case c.TotalSpent >= 50000000:
		c.LoyaltyTier = "platinum"
	case c.TotalSpent >= 10000000:
		c.LoyaltyTier = "gold"
	}

	return nil
}

func (c *CustomerProfile) UpdateAddress(address *ShippingAddress) {
	c.Address = address
}

func (c CustomerProfile) PrintProfile() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Customer Profile:")
	fmt.Printf("Name        : %s\n", c.Name)
	fmt.Printf("Loyalty Tier: %s\n", c.LoyaltyTier)
	fmt.Printf("Total Spent : %.2f\n", c.TotalSpent)
	fmt.Printf("Address     : %s, %s, %s\n", c.Address.Street, c.Address.City, c.Address.Zip)
	fmt.Println("-------------------------------------------------------")
}

func main() {
	gudangUtama := Vendor{
		Id:   "V001",
		Name: "PT. Jaya Abadi",
		Address: Address{
			Street: "Jl. Sudirman",
			City:   "Jakarta",
			Zip:    "10210",
		},
	}

	fmt.Println("==================================================")
	fmt.Println(" SISTEM GUDANG LOGISTIK E-COMMERCE ")
	fmt.Println("==================================================")

	product, err := NewProduct("P001", "Laptop", 10000000, 10, &gudangUtama)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Produc sebelum restock dan deduct stock:")
	product.PrintDetails()

	if err := product.Restock(4); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Product setelah restock:")
	product.PrintDetails()

	if err := product.DeductStock(2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Product setelah deduct stock:")
	product.PrintDetails()

	if err := product.DeductStock(20); err != nil {
		fmt.Println(err)
		return
	}

	customer := CustomerProfile{}
	customer.New("John Doe", &ShippingAddress{
		Street: "Jl. Sudirman",
		City:   "Jakarta",
		Zip:    "10210",
	})
	customer.PrintProfile()
	customer.RecordTransaction(10000000)
	customer.PrintProfile()
	customer.RecordTransaction(50000000)
	customer.PrintProfile()
}
