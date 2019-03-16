package secureJSON

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
/*
可以使用安全的json传输,其实就是添加一些干扰信息,而且只能添加头信息,真是没啥卵用
router.SecureJsonPrefix("yoawo\n") //为所有返回json添加头信息
c.SecureJSON(http.StatusOK, names)
 */
/**
 * @desc    使用 SecureJSON 来防止 json 劫持 如果给定的结构是数组值 则默认预 置 while(1)  到响应体。
 * @author Ipencil
 * @create 2019/3/16
 */

 func Secure(c *gin.Context){
		 names := []string{"lena", "austin", "foo"}
		 // Will output  :   yoawo;["lena","austin","foo"]
		 c.SecureJSON(http.StatusOK, names)
 }
