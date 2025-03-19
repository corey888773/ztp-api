package mappers

import (
	"github.com/corey888773/ztp-api/src/data"
	"github.com/corey888773/ztp-api/src/types"
)

func MapCreateProductRequest(request types.CreateProductRequest) data.Product {
	product := data.Product{
		Name:     request.Product.Name,
		Price:    request.Product.Price,
		Category: string(request.Product.Category),
		Quantity: request.Product.Quantity,
	}
	return product
}

func MapUpdateProductRequest(request types.UpdateProductRequest, id *string) data.Product {
	product := data.Product{
		ID:       *id,
		Name:     request.Product.Name,
		Price:    request.Product.Price,
		Category: string(request.Product.Category),
		Quantity: request.Product.Quantity,
	}
	return product
}
