package router

import (
	"flyblog/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.PingService{}.PingHandle)
	return r
}
