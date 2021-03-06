package show

import (
	"fmt"
	"net/http"
	"pencil/lib"

	"github.com/gin-gonic/gin"
)

/**
 * @desc   get及post解析参数,默认值,map,数组获取
 * @author Ipencil
 * @create 2019/3/14
 */
func Show(c *gin.Context) {
	user := c.MustGet("claims").(*lib.CustomClaims)
	fmt.Printf("获得jwt:%+v\n", user)
	//获取字符串,给定默认值
	firstname := c.DefaultQuery("firstname", "Guest") //填空没有默认值
	lastname := c.Query("lastname")
	queryMap := c.QueryMap("pri")
	list := c.QueryArray("list")
	c.String(http.StatusOK, "Hello %s %s map:%v list:%v", firstname, lastname, queryMap, list)
	return
}

func Posting(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	queryMap := c.PostFormMap("pri")
	list := c.PostFormArray("list")
	fmt.Println("map:", queryMap)
	fmt.Println("list:", list)
	c.JSON(200, gin.H{"message": message, "nick": nick, "map": queryMap, "list": list})
	return
}

//curl -X PUT http://localhost:8000/api/pencil/somePut
func Putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world-put", "error": nil})
	return
}

//http://localhost:8000/api/pencil/somePatch
func Patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

//http://localhost:8000/api/pencil/someDelete
func Deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

//http://localhost:8000/api/pencil/someOptions
func Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}

func Head(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello world", "error": nil})
	return
}
