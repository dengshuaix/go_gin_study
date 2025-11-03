package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/redirect-router", func(c *gin.Context) {
		c.Request.URL.Path = "/redirect"
		// 路由重定向：手动处理上下文，触发路由处理
		router.HandleContext(c)
	})

	router.GET("/redirect", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "redirect success",
		})
	})
	router.Run(":8080")
}
