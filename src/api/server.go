package api

import (
	"fmt"

	"github.com/corey888773/ztp-api/src/data"
	"github.com/corey888773/ztp-api/src/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Srv struct {
	Router          *gin.Engine
	ProductsService services.ProductsService
}

func NewServer() *Srv {
	return &Srv{
		Router: gin.Default(),
	}
}

func (s *Srv) SetupRouter() {
	products := s.Router.Group("/api/v1/products")
	products.GET("/", s.GetAllProducts)
	products.GET("/:id", ValidateParam("id", ShouldNotBeEmpty), s.GetProductById)
	products.POST("/", s.CreateProduct)
	products.PATCH("/:id", ValidateParam("id", ShouldNotBeEmpty), s.UpdateProduct)
	products.DELETE("/:id", ValidateParam("id", ShouldNotBeEmpty), s.DeleteProduct)
	products.GET("/:id/history", ValidateParam("id", ShouldNotBeEmpty), s.GetProductHistory)
}

func (s *Srv) SetupServices(mongoClient *mongo.Client) error {
	db := mongoClient.Database("ecommerce")
	fmt.Println(db.Name())

	productsCollection := db.Collection("products")
	productHistoryCollection := db.Collection("product_history")
	productRepository, err := data.NewProductRepository(mongoClient, productsCollection, productHistoryCollection)
	if err != nil {
		return err
	}

	productsService := services.NewProductsService(productRepository)
	s.ProductsService = productsService
	return nil
}

func (s *Srv) Start(httpAddress string) error {
	return s.Router.Run(httpAddress)
}
