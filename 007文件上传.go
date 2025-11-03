package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 当前路径
		pwd, _ := os.Getwd()
		fmt.Println("当前路径pwd:", pwd)
		savePath := pwd + "/static/upload/"
		fmt.Println("存储路径savePath:", savePath)

		// 保存文件，如果目录不存在就创建该目录
		if _, err := os.Stat(savePath); os.IsNotExist(err) {
			os.Mkdir(savePath, 0755)
		}
		dst := savePath + file.Filename
		c.SaveUploadedFile(file, dst)

		c.JSON(200, gin.H{
			"message": "upload success",
		})
	})
	router.Run(":8080")
}
