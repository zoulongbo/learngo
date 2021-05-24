package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	//中间件
	r.Use(func(context *gin.Context) {
		s := time.Now()
		context.Next()
		logger.Info("incoming request:",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	}, func(context *gin.Context) {
		context.Set("requestId", rand.Int())
	})

	r.GET("/ping", func(context *gin.Context) {

		h := gin.H{
			"message" : "pong",
		}

		if rqId, exist := context.Get("requestId"); exist {
			h["requestId"] = rqId
		}
		context.JSON(200, h)
	})

	r.GET("/first", func(context *gin.Context) {
		context.String(200, "%s", "I am first")
	})
	r.Run(":8080")
}
