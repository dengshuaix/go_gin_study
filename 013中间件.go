package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 计时中间件
func Timer1() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("注册顺序1")
		start := time.Now()
		c.Next() // 调用后续中间件和路由处理
		//c.Abort() // 中断后续中间件和路由处理，返回给客户端。相当于return了
		end := time.Now()
		fmt.Println("响应顺序1：", end.Sub(start))
	}
}

// 计时中间件
func Timer2() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("注册顺序2")
		start := time.Now()
		c.Next()
		end := time.Now()
		fmt.Println("响应顺序2：", end.Sub(start))
	}
}

// 计时中间件
func Timer3() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("注册顺序3")
		start := time.Now()
		c.Next()
		end := time.Now()
		fmt.Println("响应顺序3：", end.Sub(start))
	}
}

func main() {
	r := gin.Default()

	r.Use(Timer1(), Timer2()) // 注册中间件
	r.Use(Timer3())
	r.GET("/", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})

	r.Run(":8080")
}
