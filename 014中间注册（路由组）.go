package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Timer01() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		fmt.Println("耗时：", end.Sub(start))
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 创建一个路由组，路径为 "/test"
	testRouter := r.Group("/test")
	// 路由组添加中间件
	testRouter.Use(Timer01())
	// 路由组添加路由
	testRouter.GET("/", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})

	// 启动 HTTP 服务器，默认监听在 0.0.0.0:8080
	r.Run() // 等价于 r.Run(":8080")
}
