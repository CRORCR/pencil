package any

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/16
 */

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func StartPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		fmt.Println("====== Only Bind By Query String ======")
		fmt.Println(person.Name)
		fmt.Println(person.Address)
	}
	c.String(200, "Success")
}
