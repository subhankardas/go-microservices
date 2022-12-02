package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/subhankardas/go-microservices/restaurant-service/controllers"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
)

func main() {
	// Create new logger
	logger := core.NewLogger("log.json")

	// Load environment variables
	if err := env.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Create gin router with default middleware - logger and recovery (crash-free) middleware
	router := gin.Default()

	setupAPIs(router, logger)

	// Run sever on default port or PORT environment variable
	router.Run()
}

func setupAPIs(router *gin.Engine, logger core.Logger) {
	// Initialize controllers
	menuCtrl := controllers.NewMenuController(logger)

	// Setup API routes
	router.GET("/api/menu", menuCtrl.GetAllMenu)
}
