package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")
	// localhost:8080/v1/users
	usersGroup := v1.Group("/users")
	{
		usersGroup.GET("/:id", func(c *gin.Context) {
			// 处理获取用户的具体逻辑
		})
		usersGroup.POST("/", func(c *gin.Context) {

		})
		usersGroup.DELETE("/:id", func(c *gin.Context) {

		})
		usersGroup.PUT("/:id", func(c *gin.Context) {

		})
	}

	server.Run(":8080")
}
