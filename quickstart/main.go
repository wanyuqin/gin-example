package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化一个服务
	server := gin.Default()
	// 为服务定义一个API
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	// 启动服务并且监听本机的8080端口
	server.Run(":8080")
}
