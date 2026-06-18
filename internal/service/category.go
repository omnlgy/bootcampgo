package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AddCategory(ctx context.Context, pool *pgxpool.Pool, categoryName string) (string, error) {
	var id string
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"

	if err := pool.QueryRow(ctx, query, categoryName).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
