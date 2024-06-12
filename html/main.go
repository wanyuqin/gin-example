package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.LoadHTMLGlob("html/template/*")
	server.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "LoadHtmlFiles",
			"name":  "名字",
		})
	})

	server.GET("/index2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", gin.H{
			"title2": "title2",
		})
	})

	server.Run(":8080")
}
