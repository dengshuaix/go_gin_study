package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("请求路径：%s\n", c.Request.URL.Path)
		fmt.Printf("请求方法：%s\n", c.Request.Method)
		// 直接设置允许所有来源，移除域名白名单检查
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Header("Access-Control-Max-Age", "86400") // 1天缓存

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	// 给所有路由添加中间件
	r.Use(Cors())
	// 定义一个 GET 请求的路由，路径为 "/"
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	// 启动 HTTP 服务器，默认监听在 0.0.0.0:8080
	r.Run(":8080")
}
