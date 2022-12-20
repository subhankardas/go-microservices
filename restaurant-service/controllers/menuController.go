package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"github.com/subhankardas/go-microservices/restaurant-service/services"
)

type menuController struct {
	log     core.Logger
	service services.MenuService
}

// Constructor for menu controller layer.
func NewMenuController(logger core.Logger) *menuController {
	return &menuController{
		log:     logger,
		service: services.NewMenuService(logger),
	}
}

// Implementations for MenuController struct //

func (ctrl *menuController) GetAllMenu(ctx *gin.Context) {
	trxId := ctx.GetString(core.TRANSACTION_ID)

	// Get all menu, on error return server error response
	menus, err := ctrl.service.GetAllMenu(trxId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	// Success response
	ctx.JSON(http.StatusOK, menus)
}

func (ctrl *menuController) AddMenu(ctx *gin.Context) {
	trxId := ctx.GetString(core.TRANSACTION_ID)
	var menu models.Menu

	// Get request data, convert to menu, on error return bad request response
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(core.ErrInvalidRequestData))
		return
	}

	// Add new menu details, on error return server error response
	if err := ctrl.service.AddMenu(trxId, &menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
		return
	}

	// Success response
	ctx.JSON(http.StatusOK, menu)
}
