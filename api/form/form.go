package form

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @desc   : 复选框的赋值
 * @author : Ipencil
 * @date   : 2019/3/16
 */
type myForm struct {
	Colors []string `form:"colors[]"`
}

//接收数组
func FormHandler(c *gin.Context) {
	var fakeForm myForm
	err := c.ShouldBind(&fakeForm)
	if err != nil {
		fmt.Println("err", err)
	}
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
func GetByName(c *gin.Context) {
	name := c.Param("name")
	age := c.Query("age")
	fmt.Println("name:", name, "age:", age)

	// c.Query 相当于 c.Request.URL.Query().Get("age") 的简写
	//name = c.DefaultQuery("name", "hello") //设置默认值获取

	fmt.Printf("get请求参数获取成功 name[%v] age[%v]", name, age)
	c.String(http.StatusOK, "Hello [%v] [%v]", name, age)
}

type User struct {
	Name string `xml:"name" form:"name" json:"name"`
	Age  int64  `xml:"age" form:"age"  json:"age"`
}

//post参数获取 form-data格式传入-->通过PostForm获取
func Update(c *gin.Context) {
	var obj = User{}
	obj.Name = c.PostForm("name")
	obj.Age, _ = strconv.ParseInt(c.PostForm("age"), 10, 64)
	fmt.Println(obj)
	c.JSON(200, gin.H{"code": 10000, "msg": "succeed"})
}

//文件上传
func UpLoad(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Printf("upload.formfile open failed", err)
		c.JSON(200, gin.H{"code": 10001, "msg": "image open filed"})
		return
	}
	defer file.Close()
	//一般涉及到上传的接口，要避免文件明名非法字符，以及内容过大，避免用户上传一个好几十g内容
	if header.Size > 1024*1024*5 {
		fmt.Println("图片太大了")
	}
	//尽量不要用用户自己的文件名，很可能包含非法字符，导致未知的bug
	//服务端自己生成就好了
	filename := header.Filename
	fmt.Println("上传文件名:", filename)
	// 创建临时接收文件
	out, err := os.Create("copy_" + filename)
	if err != nil {
		fmt.Printf("upload.formfile failed", err)
		c.JSON(200, gin.H{"code": 10001, "msg": "image filed"})
		return
	}
	defer out.Close()
	io.Copy(out, file)
}
