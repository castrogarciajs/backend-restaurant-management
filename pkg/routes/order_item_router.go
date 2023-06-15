package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/app/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ordersItems", controllers.GetOrdersItems())
	incomingRoutes.GET("/ordersItems/:item_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemByOrder())
	incomingRoutes.POST("/ordersItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("ordersItems/:item_id", controllers.UpdateOrderItem())
}
