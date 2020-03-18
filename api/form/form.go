package form

import (
	"fmt"

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
