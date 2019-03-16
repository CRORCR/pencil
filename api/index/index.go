package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
//http://localhost:8080/index/get
/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/16
 */
 func Index(c *gin.Context){
	 c.HTML(http.StatusOK, "index.tmpl", gin.H{
		 "title": "Main website",
	 })
 }