package login

import (
	"net/http"
	"pencil/lib"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
 * @desc    登录,jwt
 * @author Ipencil
 * @create 2019/3/15
 */
func Login(c *gin.Context) {
	user := c.PostFormMap("user")
	for name, password := range user {
		if len(name) == 0 || len(password) == 0 {
			c.JSON(http.StatusOK, gin.H{"result": "用户名密码不得为空", "error": "参数异常"})
			return
		}
		if strings.EqualFold(name, "李长全") && strings.EqualFold(password, "123456") {
			//生成令牌
			createJwt(c, name, password)
			return
		}

		//检查失败,不允许登录
		c.JSON(http.StatusOK, gin.H{"result": "用户名密码不匹配", "error": "参数异常"})
	}
}

func createJwt(c *gin.Context, name, password string) {
	j := &lib.JWT{
		lib.GetSign(),
	}

	claims := lib.CustomClaims{
		Name:  name,
		Share: password, //商家id
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400*90, // 过期时间 90天
			Issuer:    "pencil",                    //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": -1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error_code": "success", "str_toke": token})
	return
}