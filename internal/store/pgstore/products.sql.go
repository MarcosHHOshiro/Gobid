// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: products.sql

package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products(
    seller_id, product_name, description, 
    baseprice, auction_end
) VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateProductParams struct {
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.SellerID,
		arg.ProductName,
		arg.Description,
		arg.Baseprice,
		arg.AuctionEnd,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT 
    seller_id,
    product_name, 
    description, 
    baseprice,
    auction_end
FROM 
    products
`

type GetAllProductsRow struct {
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

func (q *Queries) GetAllProducts(ctx context.Context) ([]GetAllProductsRow, error) {
	rows, err := q.db.Query(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllProductsRow
	for rows.Next() {
		var i GetAllProductsRow
		if err := rows.Scan(
			&i.SellerID,
			&i.ProductName,
			&i.Description,
			&i.Baseprice,
			&i.AuctionEnd,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductById = `-- name: GetProductById :one
SELECT id, seller_id, product_name, description, baseprice, auction_end, is_sold, created_at, updated_at FROM products 
WHERE id = $1
`

func (q *Queries) GetProductById(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRow(ctx, getProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.SellerID,
		&i.ProductName,
		&i.Description,
		&i.Baseprice,
		&i.AuctionEnd,
		&i.IsSold,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
