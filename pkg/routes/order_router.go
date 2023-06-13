package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/app/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:food_id", controllers.GetOrder())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("orders/:food_id", controllers.UpdateOrder())
}
