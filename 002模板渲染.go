package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// 加载 templates下所有的模板文件
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Hello, Gin!",
		})
	})

	router.Run(":8080")
}
