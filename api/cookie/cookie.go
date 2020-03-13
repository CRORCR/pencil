package cookie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/18
 */
func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil { //cookie
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, gin.H{"cookie:": cookie})
}
