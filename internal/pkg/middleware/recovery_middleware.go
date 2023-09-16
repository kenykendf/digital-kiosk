package middleware

import (
	"net/http"

	"kenykendf/digital-kiosk/internal/pkg/handler"
	"kenykendf/digital-kiosk/internal/pkg/reason"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				logrus.Error("error : ", err)
				handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
			}
		}()

		ctx.Next()
	}
}
