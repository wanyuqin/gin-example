package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	// 全局
	server.Use(GlobalMiddleware, GlobalMiddleware2)
	// 路由组
	v1 := server.Group("/v1", UserGroupMiddleWareFirst)
	// 路由
	v1.GET("/index", PathMiddleware, func(c *gin.Context) {
		// 具体业务逻辑
		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	}, func(c *gin.Context) {
		fmt.Println("after response")
	})

	v1.GET("/index2", func(c *gin.Context) {
		// 具体业务逻辑
		c.JSON(http.StatusOK, gin.H{
			"data": "ok2",
		})
	})
	server.Run(":8080")
}

func UserGroupMiddleWareFirst(c *gin.Context) {
	fmt.Println("group middle ware")
	c.Next()
}

func UserGroupMiddleWareSecond(c *gin.Context) {
	fmt.Println("second")
	c.Next()

}

func GlobalMiddleware(c *gin.Context) {
	fmt.Println(fmt.Sprintf("global middle ware request path %s\n", c.Request.RequestURI))
	c.Next()
}

func GlobalMiddleware2(c *gin.Context) {
	fmt.Println(fmt.Sprintf("global middle2 ware request path %s\n", c.Request.RequestURI))
	c.Next()
}

func PathMiddleware(c *gin.Context) {
	fmt.Println("before response")
	c.Next()
}
