package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

type RecoveryMW struct {
	log core.Logger
}

func NewRecoveryMW(log core.Logger) *RecoveryMW {
	return &RecoveryMW{log: log}
}

func (mw *RecoveryMW) RecoveryMW(ctx *gin.Context, recovered interface{}) {
	trxId := ctx.GetString(core.TRANSACTION_ID)
	if err, ok := recovered.(string); ok {
		mw.log.Errorf(trxId, "error: recovered panic, cause: %s", err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(errors.New(err)))
	} else {
		mw.log.Errorf(trxId, "error: unable to recover panic, cause: %s", recovered)
	}
	ctx.AbortWithStatus(http.StatusInternalServerError)
}
