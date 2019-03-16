package route

import (
	"pencil/api/any"
	"pencil/api/bind"
	"pencil/api/confirm"
	"pencil/api/query"
	"pencil/api/login"
	"pencil/api/show"
	"pencil/lib"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
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
	RouterGroupBind()
}

func RouterGroupLogin() {
	router := getRouter()
	router.POST("/login", login.Login)

}

/**
 * @desc   : 测试各种绑定
 * @author : Ipencil
 * @date   : 2019/3/15
 */
func RouterGroupBind() {
	router := getRouter()
	router.Any("/bind_json", bind.BandJson)//各种请求都可以支持
	router.Any("/bind_xml", bind.BandXml)//各种请求都可以支持
	router.POST("/query", query.StartPage)
	router.GET("/bookable", confirm.GetBookable)
	router.GET("/anystart", any.StartPage)
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
	_router.Use(gin.Recovery()) //中间件,异常处理
	//验证器先注册   confirm的时候,用了验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", confirm.BookableDate) //存储是以map形式存的,存储在内存中
	}
}
