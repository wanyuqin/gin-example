package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `form:"name" json:"name"`
	Age  uint   `form:"age" json:"age"`
}

func main() {
	server := gin.Default()

	v1 := server.Group("/v1")

	userGroup := v1.Group("/users")
	{
		userGroup.GET("/query", func(c *gin.Context) {
			// localhost:8080/v1/users/query?name=111
			name := c.Query("name")

			c.JSON(http.StatusOK, gin.H{
				"name": name,
			})
		})

		userGroup.POST("/struct", func(c *gin.Context) {
			user := User{}
			if err := c.ShouldBindQuery(&user); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			c.JSON(http.StatusOK, user)
		})

		userGroup.POST("/", func(c *gin.Context) {
			user := User{}
			err := c.ShouldBindJSON(&user)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
			c.JSON(http.StatusOK, user)
		})

	}
	server.Run(":8080")
}
