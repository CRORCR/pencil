package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//默认gin日志是输出控制台，可以改成输出到文件
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	gin.DefaultErrorWriter = io.MultiWriter(file)

	r := gin.New()

	r.Use(gin.Logger())
	r.GET("hello", func(context *gin.Context) {
		fmt.Println("hello world")
	})

	r.Run()
}
