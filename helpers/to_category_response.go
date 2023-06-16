package helpers

import (
	"go_restfulapi/data/response"
	"go_restfulapi/models"
)

func ToProductResponse(product models.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
	}
}
