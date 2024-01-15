package router

import (
	"gotesting/Desafio/cmd/handler"
	"gotesting/Desafio/internal/domain"
	"gotesting/Desafio/internal/product"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductRoutes(rg)
	}
}

func buildProductRoutes(r *gin.RouterGroup) {
	var productsList []domain.Product
	productsList = append(productsList, domain.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	repository := product.NewRepository(productsList)
	service := product.NewService(repository)
	handler := handler.NewHandler(service)

	productsRoutes := r.Group("/products")
	{
		productsRoutes.GET("", handler.GetProductsBySellerID)
	}

}
