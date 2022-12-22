package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/subhankardas/go-microservices/restaurant-service/controllers"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/middleware"
)

var logger core.Logger

func init() {
	// Create new logger
	logger = core.NewLogger(core.LogConfig{
		Filepath: "log.json",
		Level:    core.DebugLevel,
	})

	// Load environment variables
	if err := env.Load(); err != nil {
		logger.Errorf(core.ENV_LOAD_ERROR, "error: %s", core.UNABLE_TO_LOAD_ENV)
	}
}

func main() {
	serve()
}

func serve() {
	// Create http router with required middleware
	router := gin.New()
	router.Use(middleware.LoggingMW())
	router.Use(middleware.TimeoutMW(time.Duration(50 * time.Millisecond)))
	router.Use(gin.CustomRecovery(middleware.NewRecoveryMW(logger).RecoveryMW))

	// Setup API routes and controllers
	setupAPIs(router)

	// Run server on default port or PORT environment variable
	router.Run()
}

func setupAPIs(router *gin.Engine) {
	// Initialize controllers
	menuCtrl := controllers.NewMenuController(logger)

	// Setup API routes
	router.GET("/api/menu", menuCtrl.GetAllMenu)
	router.POST("/api/menu", menuCtrl.AddMenu)

	router.GET("/gr", func(ctx *gin.Context) { ctx.String(http.StatusOK, "num: %v", runtime.NumGoroutine()) })
}
