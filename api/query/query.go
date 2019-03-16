package query

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @desc 查询参数并绑定到结构体
 * @author Ipencil
 * @create 2019/3/16
 */
type Person struct {
	Name     string    `xml:"name" form:"name"`
	Address  string    `xml:"address" form:"address"`
	Birthday time.Time `xml:"birthday" form:"birthday" time_format:"2006-01-02 03:04:05" time_local:"8"`
}

func StartPage(c *gin.Context) {
	var person Person
	content := c.Request.Header.Get("Content-Type")
	//xml 解析
	if strings.Contains(content, "xml") {
		if err := c.BindXML(&person); err == nil {
			fmt.Printf("XML: name:%v address:%v birthday:%v\n", person.Name, person.Address, person.Birthday)
		}
	} else if strings.Contains(content, "form-data") { //json格式
		if err := c.ShouldBind(&person); err == nil {
			fmt.Printf("JSON: name:%v address:%v birthday:%v\n", person.Name, person.Address, person.Birthday)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "类型错误", "error": "content error"})
	}
	return
}