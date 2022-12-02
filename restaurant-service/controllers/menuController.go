package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/services"
)

type MenuController struct {
	log     core.Logger
	service services.MenuService
}

func NewMenuController(logger core.Logger) *MenuController {
	return &MenuController{
		log:     logger,
		service: services.NewMenuService(logger),
	}
}

// Implementations for MenuController interface //

func (ctrl *MenuController) GetAllMenu(ctx *gin.Context) {
	menu := ctrl.service.GetAllMenu()
	ctx.JSON(http.StatusOK, gin.H{
		"menu": menu,
	})
}
