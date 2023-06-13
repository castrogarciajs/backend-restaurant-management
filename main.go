package main

import (
	
	"os"
	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/platform/database"
	"github.com/sebastian009w/backend-restaurant-management/pkg/routes"
	"github.com/sebastian009w/backend-restaurant-management/pkg/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	
	r := gin.New()
	PORT := os.Getenv("PORT")

	r.Use(gin.Logger())

	if PORT == "" {
		PORT = "8000"
	}
	r.Run()
}
