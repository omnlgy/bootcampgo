package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Name       string
	SKU        string
	Price      float64
	Stock      int
	CategoryID uuid.UUID
}

type createdProduct struct {
	id string
	Product
}

func AddProduct(ctx context.Context, pool *pgxpool.Pool, product Product) (createdProduct, error) {
	var p createdProduct
	query := `
		INSERT INTO products (name, sku, price, stock, category_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, sku, price, stock, category_id
	`
	err := pool.QueryRow(ctx, query, product.Name, product.SKU, product.Price, product.Stock, product.CategoryID).Scan(
		&p.id,
		&p.Name,
		&p.SKU,
		&p.Price,
		&p.Stock,
		&p.CategoryID,
	)
	if err != nil {
		return createdProduct{}, err
	}
	return p, nil
}
