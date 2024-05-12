package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个server
	server := gin.Default()
	// 定义一个路由
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	// 启动服务
	server.Run(":8082")
}
