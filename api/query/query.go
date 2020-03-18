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
	Birthday time.Time `xml:"birthday" form:"birthday" time_format:"2006-01-02 15:04:05" time_local:"8"`
}

func StartPage(c *gin.Context) {
	var person Person
	content := c.Request.Header.Get("Content-Type")
	//xml 解析
	if strings.Contains(content, "xml") {
		if err := c.BindXML(&person); err == nil {
			c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("XML: name[%v] address[%v] birthday[%v]", person.Name, person.Address, person.Birthday), "error": nil})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "", "error": err})
		}
	} else if strings.Contains(content, "form-data") { //json格式
		if err := c.ShouldBind(&person); err == nil {
			c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("form: name[%v] address[%v] birthday[%v]", person.Name, person.Address, person.Birthday), "error": nil})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "", "error": err})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "", "error": "not found type"})
	}
	return
}
