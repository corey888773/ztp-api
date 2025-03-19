package types

import (
	"github.com/corey888773/ztp-api/src/data"
)

type GetProductsResponse struct {
	Products []data.Product `json:"products"`
}

type Category string

const (
	Electronics Category = "elektronika"
	Books       Category = "ksiazki"
	Clothing    Category = "odziez"
)

type CreateProductRequest struct {
	Product `binding:"required"`
}

type UpdateProductRequest struct {
	Product `binding:"required"`
}

type Product struct {
	Name     string   `json:"name" binding:"required,min=3,max=20,alphanum"`
	Category Category `json:"category" binding:"required,oneof=elektronika ksiazki odziez"`
	Price    float64  `json:"price" binding:"required,gt=0"`
	Quantity int      `json:"quantity" binding:"required,gte=0"`
}
