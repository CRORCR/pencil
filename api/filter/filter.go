package filter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//在中间件中使用Goroutines
//当在中间件或handler中启动新的Goroutines时，不能使用原始上下文，必须使用只读副本。
//cCp := c.Copy()

/**
 * @desc    自定义中间件
 * @author Ipencil
 * @create 2019/3/18
 */
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// Set example variable
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
	example := c.MustGet("example").(string)
	fmt.Println("filter", example)
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
