package bind

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @desc  增加标签
 * @author Ipencil
 * @create 2019/3/15
 */
type User struct {
	Name     string `form:"name" json:"name" xorm:"name" binding:"required"`
	Password string `form:"password" binding:"required" binding:"required"`
}

//绑定参数,需要加锁form,不管是get还是post请求都可以
func Band(c *gin.Context){
	user:=&User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"result":user})
}

func GetBand(c *gin.Context){
	user:=&User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"result":user})
}