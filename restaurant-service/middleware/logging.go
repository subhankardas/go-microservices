package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/utils"
)

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(core.TRANSACTION_ID, utils.NewUUID())
		ctx.Next()
	}
}
