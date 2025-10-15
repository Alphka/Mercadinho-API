package main

import (
	"log"
	"mercadinho/database"
	"mercadinho/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mercadinho/docs"
)

// @title           Mercadinho API
// @description     This is an internal API of a marketplace

// @BasePath  /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	host := "localhost:" + port

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	database.Connect()
	defer database.Pool.Close()
	database.RunMigrations()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterCustomerRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterPurchaseRoutes(router)

	router.Run(host)
}
