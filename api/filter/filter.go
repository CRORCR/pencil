package filter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//当在中间件或handler中启动新的Goroutines时，不能使用原始上下文，必须使用只读副本。
/**
 * @desc    自定义中间件
 * @author Ipencil
 * @create 2019/3/18
 */
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		fmt.Println("使用时间:", latency)
		status := c.Writer.Status()
		fmt.Println(status) //返回的状态码:200
	}
}

//返回之前走中间件
func Filter(c *gin.Context) {
	//获取以上Logger的example的值，转换成string
	example := c.MustGet("example").(string)
	fmt.Println("filter", example)
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
