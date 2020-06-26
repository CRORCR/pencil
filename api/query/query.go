package query

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

//获取body的值
//ioutil.ReadAll(response.Body)
//curl -X POST "http://localhost:8000/api/bind/body" -d "{"name":"hello"}"
func GetBody(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	c.JSON(200, gin.H{"message": "succeed"})
}

//curl -X POST "http://localhost:8000/api/bind/body2" -d "name=lcq&age=30"
//使用readall之后，会把数据都读出来，后续其他的获取参数就拿不到值了，怎么解决？
func GetBody2(c *gin.Context) {
	result, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

	//解决办法：再把body写回去就好了
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(result))
	name := c.PostForm("name")
	age := c.PostForm("age")
	fmt.Println("postForm值:", name, age)
	c.JSON(200, gin.H{"message": "succeed2"})
}
