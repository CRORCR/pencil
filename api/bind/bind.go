package bind

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/15
 */
type User struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" binding:"required"`
}

//绑定参数,需要加锁form,不管是get还是post请求都可以
func BandJson(c *gin.Context) {
	user := &User{}
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name != "李长全" || user.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": user})
}

func BandXml(c *gin.Context) {
	user := &User{}
	if err := c.ShouldBindXML(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": user})
	return
}
