package main

import (
	"fmt"
	"produkfc/cmd/product/handler"
	"produkfc/cmd/product/repository"
	"produkfc/cmd/product/resource"
	"produkfc/cmd/product/service"
	"produkfc/cmd/product/usecase"
	"produkfc/config"
	"produkfc/infrastructure/logger"
	"produkfc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDB(&cfg)

	// resource.InitRedis(&cfg)
	// resource.InitDB(&cfg)
	logger.SetupLogger()
	productRepository := repository.NewProductRepository(db, redis)
	productService := service.NewProductService(*productRepository)
	productUseCase := usecase.NewProductUsecase(*productService)
	productHandler := handler.NewProductHandler(*productUseCase)
	port := cfg.App.Port
	fmt.Print("load port = ", port)
	router := gin.Default()
	routes.SetupRoutes(router, *productHandler)
	router.Run(":" + port)
	logger.Logger.Printf("server running on port: %s", port)

}
