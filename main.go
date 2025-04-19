package main

import (
	"produkfc/cmd/config"
	"produkfc/cmd/product/resource"
	"produkfc/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	//redis := resource.InitRedis(&cfg)
	//db := resource.InitDB(&cfg)

	//resource.InitRedis(&cfg)
	resource.InitDB(&cfg)

	logger.SetupLogger()
	port := cfg.App.Port
	router := gin.Default()
	router.Run(":", port)
	logger.Logger.Printf("server running on port: %s", port)

}
