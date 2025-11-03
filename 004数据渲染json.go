package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// json响应方式1: 通过 gin.H{} 快速创建JSON响应
	r.GET("/json01", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// json响应方式2: 通过结构体创建JSON响应
	var response struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
	r.GET("/json02", func(c *gin.Context) {
		response.Status = 200
		response.Message = "pong"
		c.JSON(200, response)
	})
	r.Run(":8080")
}
