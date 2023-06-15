package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/pkg/middlewares"
	"github.com/sebastian009w/backend-restaurant-management/pkg/routes"
	"github.com/sebastian009w/backend-restaurant-management/platform/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	// Application with Gin
	r := gin.New()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	r.Use(gin.Logger())

	routes.UserRoutes(r)
	r.Use(middlewares.Authentication())

	routes.FoodRoutes(r)
	routes.MenuRoutes(r)
	routes.TableRoutes(r)
	routes.OrderRoutes(r)
	routes.OrderItemRoutes(r)
	routes.InvoiceRoutes(r)

	r.Run(":" + PORT)
}
