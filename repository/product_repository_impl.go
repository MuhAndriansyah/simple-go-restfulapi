package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_restfulapi/helpers"
	"go_restfulapi/models"
	"time"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product models.Product) models.Product {
	SQL := "INSERT INTO products (name, price, description) VALUES (?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Description)
	helpers.PanicIfError(err)

	productId, err := result.LastInsertId()
	helpers.PanicIfError(err)

	product.Id = int(productId)
	product.UpdatedAt = time.Now()
	product.CreatedAt = time.Now()

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product models.Product) models.Product {
	SQL := "UPDATE products SET name = ?, price = ?, description = ?, updated_at = ? WHERE id = ?"

	currentTime := time.Now()

	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Description, currentTime, product.Id)
	helpers.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product models.Product) {
	SQL := "DELETE FROM products WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helpers.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (models.Product, error) {
	SQL := "SELECT id, name, price, description, updated_at, created_at FROM products WHERE id = ?"

	rows, err := tx.QueryContext(ctx, SQL, productId)
	helpers.PanicIfError(err)
	defer rows.Close()

	product := models.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.UpdatedAt, &product.CreatedAt)
		helpers.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("product tidak ditemukan")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []models.Product {
	SQL := "SELECT id, name, price, description, updated_at, created_at FROM products"

	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)

	var products []models.Product

	for rows.Next() {
		product := models.Product{}

		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.UpdatedAt, &product.CreatedAt)
		helpers.PanicIfError(err)

		products = append(products, product)
	}

	return products
}
