package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 添加中间件（可选）
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 定义一个 GET 请求的路由，路径为 "/ping"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 定义一个带参数的路由
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	// 启动 HTTP 服务器，默认监听在 0.0.0.0:8080
	r.Run() // 等价于 r.Run(":8080")
}
