package test

import (
	"fmt"
	"testing"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/18
 */
func filter(t *testing.T) {
	t.SkipNow()
	t.Run("band_json", filterPrint)
}

//json客户端发送数据
func filterPrint(t *testing.T) {
	t.SkipNow()
	/*get 请求*/
	url := "http://localhost:8080/filter" //填空没有默认值
	result := queryGet(t, url)
	fmt.Println(result)
}
