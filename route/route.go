package route

import (
	"net/http"
	"pencil/api/any"
	"pencil/api/bind"
	"pencil/api/confirm"
	"pencil/api/filter"
	"pencil/api/form"
	"pencil/api/index"
	"pencil/api/login"
	"pencil/api/query"
	"pencil/api/secureJSON"
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
	RouterGroupIndex() //模板渲染
	RouterGroupFilter()
}

func RouterGroupLogin() {
	router := getRouter()
	router.POST("/login", login.Login)
}
//中间件
func RouterGroupFilter() {
	router := getRouter()
	router.Use(filter.Logger()) //加载中间件
	router.GET("/filter",filter.Filter)
}

func RouterGroupIndex() {
	router := getRouter()
	router.GET("/index/get", index.Index) //模板渲染
	router.GET("/redir", index.Redirct)   //重定向
	//路由重定向
	router.GET("/redirhand",
		func(c *gin.Context) {
			c.Request.URL.Path = "/hello"
			router.HandleContext(c)
		})
	router.GET("/hello", index.Hello)

}

/**
 * @desc   : 测试各种绑定
 * @author : Ipencil
 * @date   : 2019/3/15
 */
func RouterGroupBind() {
	router := getRouter()
	router.Any("/bind_json", bind.BandJson) //各种请求都可以支持,不支持多次序列化
	router.POST("/bandbind", bind.BandJsonBind) //各种请求都可以支持,并且可以支持多次使用,多个if else
	router.Any("/bind_xml", bind.BandXml)   //各种请求都可以支持
	router.POST("/query", query.StartPage)
	router.GET("/bookable", confirm.GetBookable)
	router.GET("/anystart", any.StartPage)
	router.Any("/forms", form.FormHandler)
	router.SecureJsonPrefix("yoawo\n") //为所有返回json添加头信息
	router.GET("/someJSON", secureJSON.Secure)
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
	//首先需要是生成一个Engine 这是gin的核心 默认带有Logger 和 Recovery 两个中间件
	_router = gin.Default()
	_router.Use(gin.Recovery()) //中间件,异常处理
	//验证器先注册   confirm的时候,用了验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", confirm.BookableDate) //存储是以map形式存的,存储在内存中
	}

	//这些目录下资源是可以随时更新，而不用重新启动程序
	_router.Static("/assets", "./assets")
	// StaticFile 是加载单个文件 StaticFS 是加载一个完整的目录资源
	_router.StaticFS("/more_static", http.Dir("my_file_system"))
	_router.StaticFile("/pencil.go", "K:/workspace/src/pencil/pencil.go")
	//模板渲染
	//_router.LoadHTMLGlob("templates/*")
	_router.LoadHTMLGlob("templates/**/*")
}

/*
http://localhost:8080/assets/doc.html  静态文件内容,可以随意访问

 */
