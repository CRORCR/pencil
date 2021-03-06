package route

import (
	"net/http"
	"pencil/lib"

	"pencil/api/bind"
	"pencil/api/confirm"
	"pencil/api/cookie"
	"pencil/api/filter"
	"pencil/api/form"
	"pencil/api/index"
	"pencil/api/login"
	"pencil/api/query"
	"pencil/api/show"
	"strings"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/14
 */
const routerV1 = "/api/pencil"
const routerV2 = "/api/bind"

func GetInitRouter() *gin.Engine {
	//首先需要是生成一个Engine 这是gin的核心 默认带有Logger 和 Recovery 两个中间件
	router := gin.Default()
	//中间件 Use设置中间件
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*", "lang", "json-web-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	v1 := router.Group(routerV1)
	{
		bindMethod(v1, "POST", "/login", login.Login) //接收map
		bindMethod(v1, "GET", "/cook", cookie.Cookie)

		v1.PUT("/somePut", show.Putting)
		v1.PATCH("/somePatch", show.Patching)
		v1.DELETE("/someDelete", show.Deleting)
		v1.HEAD("/someHead", show.Head) //不好使
		v1.OPTIONS("/someOptions", show.Options)

		v1.Use(lib.JWTAuth())
		bindMethod(v1, "GET", "/show", show.Show)
		bindMethod(v1, "POST", "/uploada", show.UploadAll)
		bindMethod(v1, "POST", "/upload", show.UploadOne)
		bindMethod(v1, "POST", "/somePost", show.Posting)
	}

	//验证器先注册   confirm的时候,用了验证器，现在没有成功，后面再看看吧
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", confirm.BookableDate) //存储是以map形式存的,存储在内存中
	}

	//测试各种绑定
	v2 := router.Group(routerV2)
	{
		v2.Any("/bind_json", bind.BandJson)     //各种请求都可以支持,不支持多次序列化
		v2.POST("/bandbind", bind.BandJsonBind) //各种请求都可以支持,并且可以支持多次使用,多个if else
		v2.Any("/query", query.StartPage)
		v2.POST("/body", query.GetBody)
		v2.POST("/body2", query.GetBody2)
		v2.GET("/bookable", confirm.GetBookable)
		v2.Any("/forms", form.FormHandler) //接收数组
		router.SecureJsonPrefix("yoawo\n") //为所有返回json添加头信息
	}

	//此规则能够匹配/user/lcq/30这种格式，但不能匹配/user/李长全/30 不支持中文，而且也不能为空，否则404
	router.GET("/users/:name/:age", form.GetNameAndAge)
	router.POST("/users/update", form.Update)
	router.POST("/users/upload", form.UpLoad)

	//加载自定义中间件
	router.Use(filter.Logger())
	router.GET("/filter", filter.Filter)

	//模板渲染
	router.GET("/index/get", index.Index) //模板渲染
	router.GET("/redir", index.Redirct)   //重定向到淘宝
	//服务端路由重定向
	router.GET("/redirhand",
		func(c *gin.Context) {
			c.Request.URL.Path = "/hello"
			router.HandleContext(c)
		})
	router.GET("/hello", index.Hello) //内部重定向到/hello路由去处理

	//这些目录下资源是可以随时更新，而不用重新启动程序
	router.Static("/assets", "./assets")
	// StaticFile 是加载单个文件 StaticFS 是加载一个完整的目录资源
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/pencil.go", "K:/workspace/src/pencil/pencil.go")
	//模板渲染
	//_router.LoadHTMLGlob("templates/*")
	router.LoadHTMLGlob("templates/**/*")

	return router
}

func bindMethod(group *gin.RouterGroup, method, path string, handler gin.HandlerFunc) {
	method = strings.ToUpper(method)
	switch method {
	case "GET":
		group.GET(path, handler)
	case "POST":
		group.POST(path, handler)
	case "PUT":
		group.PUT(path, handler)
	default:
		panic("not find method")
	}
}

/*
http://localhost:8000/assets/doc.html  静态文件内容,可以随意访问
*/
