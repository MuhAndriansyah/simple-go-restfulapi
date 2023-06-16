package services

import (
	"context"
	"database/sql"
	"go_restfulapi/data/request"
	"go_restfulapi/data/response"
	"go_restfulapi/exception"
	"go_restfulapi/helpers"
	"go_restfulapi/models"
	"go_restfulapi/repository"
)

type ProductServiceImpl struct {
	DB                *sql.DB
	ProductRepository repository.ProductRepository
}

func NewProductService(db *sql.DB, productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		DB:                db,
		ProductRepository: productRepository,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
	}

	product = service.ProductRepository.Create(ctx, tx, product)

	return helpers.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError("product tidak ditemukan"))
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Description = request.Description

	product = service.ProductRepository.Update(ctx, tx, product)

	return helpers.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError("product tidak ditemukan"))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) response.ProductResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError("product tidak ditemukan"))
	}

	return helpers.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productResponses []response.ProductResponse

	for _, product := range products {
		productResponses = append(productResponses, helpers.ToProductResponse(product))
	}

	return productResponses
}
