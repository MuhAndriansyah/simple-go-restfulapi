package services

import (
	"context"
	"go_restfulapi/data/request"
	"go_restfulapi/data/response"
)

type ProductService interface {
	Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) response.ProductResponse
	FindAll(ctx context.Context) []response.ProductResponse
}
