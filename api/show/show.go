package show

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/14
 */
func Show(c *gin.Context) {
	name := c.Param("name")
	name2 := c.Param("name2")
	fmt.Println("name", name, "name2", name2)
	c.String(http.StatusOK, "Hello %v", name)
	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("Hello:%v", name), "error": nil})
	return
}

func Show2(c *gin.Context) {
	//获取字符串,给定默认值
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	return
}

func Posting(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{"status": "posted", "message": message, "nick": nick})
	return
}

func Putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

func Deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

func Patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

func Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

func Head(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}
