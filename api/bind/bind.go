package bind

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/**
 * @desc    json
 * @author Ipencil
 * @create 2019/3/15
 */
type User struct {
	Name     string `xml:"name" form:"name" json:"name"`
	Password string `xml:"password" form:"password"  json:"password"`
}

//ShouldBind 绑定参数 不管是get还是post请求都可以,可以多次调用，多次调用有什么用？
func BandJson(c *gin.Context) {
	user := &User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("user1:", user)

	user2 := &User{}
	if err := c.ShouldBindJSON(user2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("user2:", user2)

	if user.Name != "李长全" || user.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": user})
}

//post请求,可多次绑定  可以不使用这种方式绑定
//c.ShouldBindBodyWith会在绑定之前将body存储到上下文中。 这会对性能造成轻微影响，
// 如果调用一次就能完成绑定的话，那就不要用这个方法。
//只有某些格式需要此功能 ，如JSON、XML、MsgPack、ProtoBuf。 对于其他格式，
// 如Query、Form、FormPost、FormMultipart可以多次调用c.ShouldBind()而不会造成任任何性能损失
func BandJsonBind(c *gin.Context) {
	user := &User{}
	if err := c.ShouldBindBodyWith(user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name != "李长全" || user.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": user})
}
