package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/data"
)

func GetAllMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.GetAllMenu())
}
