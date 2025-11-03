package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// xml 响应方式1: 通过 gin.H{} 快速创建xml响应
	r.GET("/xml01", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "pong",
			"id":      "001",
		})
	})

	// xml响应方式2: 通过结构体创建xml响应
	type Message struct {
		Message string `xml:"message"`
		Status  string `xml:"status"`
		ID      int    `xml:"id"`
	}
	var response Message
	r.GET("/xml02", func(c *gin.Context) {
		response.Status = "200"
		response.Message = "pong"
		response.ID = 1001
		c.XML(http.StatusOK, response)
	})
	r.Run(":8080")
}
