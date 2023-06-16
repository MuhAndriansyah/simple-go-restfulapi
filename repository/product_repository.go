package repository

import (
	"context"
	"database/sql"
	"go_restfulapi/models"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product models.Product) models.Product
	Update(ctx context.Context, tx *sql.Tx, product models.Product) models.Product
	Delete(ctx context.Context, tx *sql.Tx, product models.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (models.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []models.Product
}
