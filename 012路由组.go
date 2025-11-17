package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//URL前缀的路由划分为一个路由组。习惯性一对{}包裹
	// 创建一个默认的路由引擎
	r := gin.Default()
	routerGroup1 := r.Group("/api")
	routerGroup1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api",
		})
	})
	routerGroup2 := r.Group("/api2")
	routerGroup2.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api2",
		})
	})
	r.Run()
}
