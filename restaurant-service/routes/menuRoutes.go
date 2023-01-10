package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/controllers"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/data"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"github.com/subhankardas/go-microservices/restaurant-service/services"
)

func SetupMenuRoutes(router *gin.Engine, config *models.Config, logger core.Logger) {
	// Initialize required dependencies for controller
	db := core.NewDatabase(config, logger)
	menuData := data.NewMenuData(config, logger, db)
	menuService := services.NewMenuService(config, logger, menuData)

	// Initialize controllers
	menuCtrl := controllers.NewMenuController(config, logger, menuService)

	// Setup API routes
	router.GET("/api/menu", menuCtrl.GetAllMenu)
	router.POST("/api/menu", menuCtrl.AddMenu)
	router.PUT("/api/menu/:id", menuCtrl.UpdateMenu)
	router.DELETE("/api/menu/:id", menuCtrl.DeleteMenu)
}
