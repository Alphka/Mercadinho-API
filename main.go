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

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		database.Connect()
		defer database.Pool.Close()
		database.RunMigrations(migrationsFS)
		return
	}

	host := ""
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	_, isVercel := os.LookupEnv("VERCEL")

	if isVercel {
		vercelTargetEnv := os.Getenv("VERCEL_TARGET_ENV")
		isProduction := vercelTargetEnv == "production" || vercelTargetEnv == "preview"

		if isProduction {
			gin.SetMode(gin.ReleaseMode)
		}

		vercelDomain := os.Getenv("VERCEL_URL")

		if vercelDomain == "" {
			log.Fatalf("VERCEL_URL environment variable not set")
		}

		docs.SwaggerInfo.Host = vercelDomain
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		host = os.Getenv("HOST")

		if host == "" {
			host = "localhost"
		}

		docs.SwaggerInfo.Host = host + ":" + port
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	docs.SwaggerInfo.Version = "1.0"

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

	database.Connect()
	defer database.Pool.Close()

	if !isVercel {
		database.RunMigrations(migrationsFS)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies(nil)

	router.HEAD("/", func(context *gin.Context) {
		context.Status(200)
	})

	router.GET("/", func(context *gin.Context) {
		context.Redirect(302, "/swagger/index.html")
	})

	router.GET("/swagger", func(context *gin.Context) {
		context.Redirect(301, "/swagger/index.html")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterCustomerRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterPurchaseRoutes(router)
	routes.RegisterHeadRoutes(router)

	router.Run(host + ":" + port)
}
