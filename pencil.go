package main

import (
	"pencil/db"
	"pencil/route"
)

const (
	YOAWOENV = "localhost"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/14
 */
func main() {
	db.InitConfig(YOAWOENV)
	route.InitRoute()
	route.GroupRouter()
	route.GlobalRout(":8080") //测试端口
}
