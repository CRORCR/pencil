package test

import (
	"fmt"
	"pencil/api/bind"
	"pencil/api/query"
	"testing"
	"time"
)
/*

any请求方式,支持如下几种:GET \ POST \ PUT \ PATCH \ HEAD \ OPTIONS \ DELETE \ CONNECT \ TRACE
xml格式数据和json格式数据传递的时候,可以通过 Content-Type 检查,xml格式的时候,结构体必须有xml标识,json必须有form标识

*/

/**
 * @desc   绑定json \ xml 测试     验证器测试
 * @author Ipencil
 * @create 2019/3/16
 */
func band(t *testing.T) {
	t.SkipNow()
	t.Run("band_json", band_json) //json解析的接口,get和post使用同一个,都支持,使用Any请求方式
	t.Run("band_json_post", band_json_post) //支持get和post方式
	t.Run("band_xml", band_xml) //支持get和post方式
	t.Run("books", books) //自定义验证器
	t.Run("query", queryJson)  //支持json和xml两种格式
	t.Run("query", queryXML)  //支持json和xml两种格式
}

//json客户端发送数据
func queryJson(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/query"
	params := map[string]string{
		"name":     "李长全",
		"address": "123",
		"birthday": "1992-08-25  12:12:12",
	}
	send := postSend(url, params)
	fmt.Println(send)
}

//json客户端发送数据
func queryXML(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/query"
	per :=query.Person{}
	per.Name = "lcq"
	per.Address = "123"
	per.Birthday=time.Now()
	send := postSendCopy(url, per)
	fmt.Println(send)
}

//json客户端发送数据
func band_json(t *testing.T) {
	t.SkipNow()
	/*get 请求*/
	url := "http://localhost:8080/bind_json?name=李长全&password=123" //填空没有默认值
	result := queryGet(t, url)
	fmt.Println(result)
}

/*post 请求*/
func band_json_post(t *testing.T){
	t.SkipNow()
	url := "http://localhost:8080/bind_json"
	params := map[string]string{
		"name":     "李长全",
		"password": "123", //这种形式也算有值,不会填充默认值
	}
	send := postSend(url, params)
	fmt.Println(send)
}

//xml 客户端发送数据
func band_xml(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/bind_xml"
	user := bind.User{}
	user.Name = "lcq"
	user.Password = "123"
	send := postSendCopy(url, user)
	fmt.Println(send)
}

//自定义验证器
func books(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/bookable?check_in=2019-04-16&check_out=2019-04-17" //填空没有默认值
	result := queryGet(t, url)
	fmt.Println(result)
}
