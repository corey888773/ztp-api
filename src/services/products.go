package services

import (
	"github.com/corey888773/ztp-api/src/data"
	"github.com/corey888773/ztp-api/src/mappers"
	"github.com/corey888773/ztp-api/src/types"
)

type ProductsService interface {
	GetAllProducts() ([]data.Product, error)
	GetProductById(id string) (*data.Product, error)
	CreateProduct(product types.CreateProductRequest) error
	UpdateProduct(product types.UpdateProductRequest, id string) error
	DeleteProduct(id string) error
	GetProductHistory(id string) ([]data.ProductChange, error)
}

type productsService struct {
	productRepository data.ProductRepository
}

func NewProductsService(repository data.ProductRepository) ProductsService {
	return &productsService{
		productRepository: repository,
	}
}

func (s *productsService) GetProductHistory(id string) ([]data.ProductChange, error) {
	return s.productRepository.GetProductHistory(id)
}

func (s *productsService) GetAllProducts() ([]data.Product, error) {
	return s.productRepository.FindAll()
}

func (s *productsService) GetProductById(id string) (*data.Product, error) {
	return s.productRepository.FindById(id)
}

func (s *productsService) CreateProduct(product types.CreateProductRequest) error {
	return s.productRepository.Create(mappers.MapCreateProductRequest(product))
}

func (s *productsService) UpdateProduct(product types.UpdateProductRequest, id string) error {
	return s.productRepository.Update(mappers.MapUpdateProductRequest(product, &id))
}

func (s *productsService) DeleteProduct(id string) error {
	return s.productRepository.Delete(id)
}
