package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

func TimeoutMW(timeout time.Duration) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		done := make(chan bool)

		// Handle next http handler in routine
		go func() {
			ctx.Next()
			done <- true
		}()

		select {
		// Timeout request and abort next handlers
		case <-time.After(timeout):
			ctx.JSON(http.StatusRequestTimeout, models.NewErrorResponse(errors.New("request timeout")))
			ctx.Abort()
		case <-done:
		}
	}
}
