package service

import "github.com/gin-gonic/gin"

type PingService struct {
}

func (p PingService) PingHandle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
