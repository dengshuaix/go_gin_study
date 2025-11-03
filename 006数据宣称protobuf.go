package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/protobuf", func(c *gin.Context) {

		reps := []int64{int64(1), int64(2), int64(3)}
		label := "test"
		// protobuf 的具体定义
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
