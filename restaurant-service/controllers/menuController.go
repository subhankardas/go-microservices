package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
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

// Implementations for MenuController struct //

func (ctrl *MenuController) GetAllMenu(ctx *gin.Context) {
	trxId := ctx.GetString(core.TRANSACTION_ID)
	menus, err := ctrl.service.GetAllMenu(trxId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

func (ctrl *MenuController) AddMenu(ctx *gin.Context) {
	trxId := ctx.GetString(core.TRANSACTION_ID)
	menu := models.Menu{}
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if err := ctrl.service.AddMenu(trxId, &menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}
