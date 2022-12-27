package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/controllers"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/middleware"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

var logger core.Logger
var config *models.Config

func init() {
	// Load config based on profile i.e. config filename passed as argument
	config = core.LoadConfig("./configs/", os.Args[1], "yml")

	// Create new logger
	logger = core.NewLogger(config.Log)
}

func main() {
	serve()
}

func serve() {
	// Create http router with required middleware
	router := gin.New()
	router.Use(middleware.LoggingMW())
	router.Use(middleware.TimeoutMW(config.Server.RequestTimeoutDuration))
	router.Use(gin.CustomRecovery(middleware.NewRecoveryMW(logger).RecoveryMW))

	// Setup API routes and controllers
	setupAPIs(router)

	// Run server on default/given port or PORT environment variable
	if err := router.Run(config.Server.Port); err != nil {
		logger.Fatalf(core.SERVER_ERROR, "error: %s", core.UNABLE_TO_RUN_SERVER)
	}
}

func setupAPIs(router *gin.Engine) {
	// Initialize controllers
	menuCtrl := controllers.NewMenuController(config, logger)

	// Setup API routes
	router.GET("/api/menu", menuCtrl.GetAllMenu)
	router.POST("/api/menu", menuCtrl.AddMenu)
}
