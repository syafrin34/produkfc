package main

import (
	"fmt"
	"produkfc/cmd/product/resource"
	"produkfc/config"
	"produkfc/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	//redis := resource.InitRedis(&cfg)
	//db := resource.InitDB(&cfg)

	resource.InitRedis(&cfg)
	resource.InitDB(&cfg)
	logger.SetupLogger()
	port := cfg.App.Port
	fmt.Print("load port = ", port)
	router := gin.Default()
	router.Run(":" + port)
	logger.Logger.Printf("server running on port: %s", port)

}
