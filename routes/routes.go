package routes

import (
	"produkfc/cmd/product/handler"
	"produkfc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productHandler handler.ProductHandler) {
	router.Use(middleware.RequestLogger())
	router.POST("/v1/product_category", productHandler.ProductCategoryManagement)
	router.POST("/v1/product", productHandler.ProductManagement)

	router.GET("/v1/product/:id", productHandler.GetProductInfo)
	router.GET("/v1/product_category/:id", productHandler.GetProductCategoryInfo)
	router.GET("/v1/product/search", productHandler.SearchProduct)
}
