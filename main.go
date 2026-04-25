package main

import (
	"embed"
	"log"
	"mercadinho/database"
	"mercadinho/helpers"
	"mercadinho/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mercadinho/docs"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// @title           Mercadinho API
// @description     This is an internal API of a marketplace

// @BasePath  /api/v1
func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.New(os.Stderr, "", 1).Println("Error loading .env file")
		}
	}

	host := ""
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	_, isProduction := os.LookupEnv("VERCEL")

	if isProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		host = os.Getenv("HOST")

		if host == "" {
			host = "localhost"
		}

		docs.SwaggerInfo.Host = host + ":" + port
		docs.SwaggerInfo.Version = "1.0"
		// docs.SwaggerInfo.Schemes = []string{"http", "https"}

		if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
			validate.RegisterValidation("date", func(fieldLevel validator.FieldLevel) bool {
				_, err := helpers.ParseDate(fieldLevel.Field().String())
				return err == nil
			})

			validate.RegisterValidation("price", func(fieldLevel validator.FieldLevel) bool {
				price, ok := fieldLevel.Field().Interface().(float64)
				if !ok {
					return false
				}

				if price <= 0 {
					return false
				}

				scaledPrice := price * 100
				if scaledPrice != float64(int64(scaledPrice)) {
					return false
				}

				return true
			})

			validate.RegisterValidation("phone", func(fieldLevel validator.FieldLevel) bool {
				_, err := helpers.NormalizePhone(fieldLevel.Field().String())
				return err == nil
			})
		}
	}

	database.Connect()
	defer database.Pool.Close()
	database.RunMigrations(migrationsFS)

	router := gin.Default()
	router.Use(cors.Default())

	if !isProduction {
		router.SetTrustedProxies(nil)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	routes.RegisterCustomerRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterPurchaseRoutes(router)

	router.Run(host + ":" + port)
}
