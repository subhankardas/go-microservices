package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/subhankardas/go-microservices/restaurant-service/controllers"
)

func main() {
	// Load environment variables
	if err := env.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Create gin router with default middleware - logger and recovery (crash-free) middleware
	router := gin.New()

	// Setup routers
	router.GET("/api/menu", controllers.GetAllMenu)

	// Run sever on default port or PORT environment variable
	router.Run()
}
