package route

import (
	"pencil/api/login"
	"pencil/api/show"
	"pencil/lib"

	"github.com/gin-gonic/gin"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/14
 */
var _router *gin.Engine

func getRouter() *gin.Engine {
	return _router
}

func GlobalRout(strPort string) {
	rout := getRouter()
	rout.Run(strPort)
}

func GroupRouter() {
	RouterGroupHello("pencil")
	RouterGroupLogin()
}

func RouterGroupLogin() {
	router := getRouter()
	router.POST("/login", login.Login)

}

func RouterGroupHello(name string) {
	//engine:=gin.Default() 默认初始化gin,然后去创建组函数
	//engine.Group()
	//Default返回一个引擎实例，其中已经附加了日志记录器和恢复中间件
	router := getRouter().Group(name)
	router.Use(lib.JWTAuth())
	router.GET("/show", show.Show)
	router.POST("/somePost", show.Posting)
	router.PUT("/somePut", show.Putting)
	router.PATCH("/somePatch", show.Patching)
	router.DELETE("/someDelete", show.Deleting)
	router.HEAD("/someHead", show.Head)
	router.OPTIONS("/someOptions", show.Options)
	router.POST("/upload", show.UploadOne)
	router.POST("/uploada", show.UploadAll)
}

func InitRoute() {
	_router = gin.Default()
	_router.Use(gin.Logger())   //添加日志,默认控制台
	_router.Use(gin.Recovery()) //中间件,异常处理
}
