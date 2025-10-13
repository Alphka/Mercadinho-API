package routes

import (
	"mercadinho/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		customers := v1.Group("/purchases")
		{
			customers.GET("", controllers.GetPurchases)
			customers.GET("/:id", controllers.GetPurchase)
		}
	}
}
