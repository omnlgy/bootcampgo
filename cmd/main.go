package main

import (
	"context"
	"fmt"

	"example.com/internal/service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	constring := "postgres://admin:admin123@localhost:5432/day6"

	dbpool, err := pgxpool.New(context.Background(), constring)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer dbpool.Close()

	err = dbpool.Ping(context.Background())
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	fmt.Println("Connected to database")

	// Test category service
	categoryID, err := service.AddCategory(context.Background(), dbpool, "Test Category")
	if err != nil {
		fmt.Println("Error adding category:", err)
		return
	}
	fmt.Println("Added category with ID:", categoryID)

	product := service.Product{
		Name:       "Test Product",
		SKU:        "TEST123",
		Price:      100.00,
		Stock:      10,
		CategoryID: uuid.MustParse(categoryID),
	}

	createdProduct, err := service.AddProduct(context.Background(), dbpool, product)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return
	}
	fmt.Println("Added product:", createdProduct)
}
