package routes

import (
	"produkfc/cmd/product/handler"
	"produkfc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productHandler handler.ProductHandler) {
	router.Use(middleware.REquestLogger())
	router.POST("/v1/product", productHandler.ProductCategoryManagement)
}
