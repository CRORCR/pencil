package index

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//http://localhost:8000/index/get
/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/16
 */
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func Redirct(c *gin.Context) {
	query := struct {
		Name string `json:"name" form:"name"`
	}{}
	c.ShouldBind(&query)
	fmt.Println(query.Name) //重新向之后，还可以获取参数
	c.Redirect(http.StatusMovedPermanently, "http://www.taobao.com/")
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello world!!!")
}
