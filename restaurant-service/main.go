package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/subhankardas/go-microservices/restaurant-service/controllers"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/middleware"
)

func main() {
	// Create new logger
	logger := core.NewLogger(core.LogConfig{
		Filepath: "log.json",
		Level:    core.DebugLevel,
	})

	// Load environment variables
	if err := env.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Create gin router
	router := gin.New()
	router.Use(middleware.Logging(logger), gin.Recovery())

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
