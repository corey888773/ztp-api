package api

import (
	"net/http"

	"github.com/corey888773/ztp-api/src/errors"
	"github.com/corey888773/ztp-api/src/types"
	"github.com/gin-gonic/gin"
)

func (s *Srv) GetAllProducts(ctx *gin.Context) {
	products, err := s.ProductsService.GetAllProducts()
	if err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (s *Srv) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := s.ProductsService.GetProductById(id)
	if err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (s *Srv) CreateProduct(ctx *gin.Context) {
	var createProductRequest types.CreateProductRequest
	if err := ctx.ShouldBindJSON(&createProductRequest); err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	if err := s.ProductsService.CreateProduct(createProductRequest); err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse())
}

func (s *Srv) UpdateProduct(ctx *gin.Context) {
	var updateProductRequest types.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&updateProductRequest); err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	id := ctx.Param("id")
	if err := s.ProductsService.UpdateProduct(updateProductRequest, id); err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, SuccessResponse())
}

func (s *Srv) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := s.ProductsService.DeleteProduct(id); err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, SuccessResponse())
}

func (s *Srv) GetProductHistory(ctx *gin.Context) {
	id := ctx.Param("id")
	productHistory, err := s.ProductsService.GetProductHistory(id)
	if err != nil {
		custom_errors.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productHistory)
}
