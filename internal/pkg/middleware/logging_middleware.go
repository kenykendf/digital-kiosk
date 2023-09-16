package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()               //starting time request
		ctx.Next()                            //process request
		endTime := time.Now()                 //end time request
		latencyTime := endTime.Sub(startTime) //calculate execution time
		requestMehod := ctx.Request.Method    //request method
		reqURI := ctx.Request.RequestURI      //request route
		statusCode := ctx.Writer.Status()     //status code
		clientIp := ctx.ClientIP()            //client ip

		log.WithFields(log.Fields{
			"latency_time":   latencyTime,
			"request_method": requestMehod,
			"req_uri":        reqURI,
			"status_code":    statusCode,
			"client_ip":      clientIp,
		}).Info("http request")

		ctx.Next()
	}
}
