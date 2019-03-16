package test

import (
	"fmt"
	"pencil/api/bind"
	"testing"
)

/**
 * @desc   绑定json \ xml 测试     验证器测试
 * @author Ipencil
 * @create 2019/3/16
 */
func band(t *testing.T) {
	t.Run("band_xml", band_xml)
	t.Run("band_json", band_json)
	t.Run("books", books)
}

//json客户端发送数据
func band_json(t *testing.T) {
	t.SkipNow()
	/*get 请求*/
	url := "http://localhost:8080/bind_json_get?name=李长全&password=123" //填空没有默认值
	result := queryGet(t, url)
	fmt.Println(result)
	/*post 请求*/

	//url := "http://localhost:8080/bind_json_post"
	//params := map[string]string{
	//	"name":     "李长全",
	//	"password": "123", //这种形式也算有值,不会填充默认值
	//}
	//send := postSend(url, params)
	//fmt.Println(send)
}

//xml 客户端发送数据
func band_xml(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/bind_xml_post"
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
