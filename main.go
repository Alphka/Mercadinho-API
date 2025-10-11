package main

import (
	"mercadinho/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mercadinho/docs"
)

// @title           Mercadinho API
// @description     This is an internal API of a marketplace

// @BasePath  /api/v1
func main() {
	var host = "localhost:8080"

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterCustomerRoutes(router)

	router.Run(host)
}
