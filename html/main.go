package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	//server.LoadHTMLFiles("html/template/index.html", "html/template/index2.html")
	server.LoadHTMLGlob("html/template/*")
	server.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello world",
		})
	})
	server.GET("/index2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", gin.H{
			"title2": "Hello world2",
		})
	})
	server.Run(":8080")
}
