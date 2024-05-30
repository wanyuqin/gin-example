package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")
	v1.Use(func(c *gin.Context) {
		fmt.Println("v1 ")
		c.Next()
	})

	userGroup := v1.Group("/users")
	userGroup.Use(UserGroupMiddleWareFirst, UserGroupMiddleWareSecond)
	userGroup.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	server.Run(":8080")
}

func UserGroupMiddleWareFirst(c *gin.Context) {
	fmt.Println("fist")
	c.Next()
}

func UserGroupMiddleWareSecond(c *gin.Context) {
	fmt.Println("second")
	c.Next()

}
