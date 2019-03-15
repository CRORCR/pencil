package route

import (
	"pencil/api/show"

	"github.com/gin-gonic/gin"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/14
 */
var g_Router *gin.Engine

func getRouter() *gin.Engine {
	return g_Router
}

func GlobalRout(strPort string) {
	rout := getRouter()
	rout.Run(strPort)
}

func GroupRouter() {
	RouterGroupHello("pencil")
}

func RouterGroupHello(name string) {
	router := getRouter().Group(name)
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

func InitRout() {
	g_Router = gin.Default()
	//g_Router.Use(cors.New(cors.Config{
	//	AllowOriginFunc:  func(origin string) bool { return true },
	//	AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"},
	//	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))
}
