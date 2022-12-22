package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/utils"
)

func LoggingMW() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Set transaction ID used for logging and audit
		ctx.Set(core.TRANSACTION_ID, utils.NewUUID())
		ctx.Next()
	}
}
