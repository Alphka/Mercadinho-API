package routes

import (
	"mercadinho/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		purchases := v1.Group("/purchases")
		{
			purchases.GET("", controllers.GetPurchases)
			purchases.GET("/:id", controllers.GetPurchase)
			purchases.PUT("/:id", controllers.UpdatePurchase)
			purchases.DELETE("/:id", controllers.DeletePurchase)
		}
	}
}
