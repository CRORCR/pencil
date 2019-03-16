package test

import (
	"fmt"
	"net/http"
	"testing"
)

/**
 * @desc   get post delete 等请求类型
 * @author Ipencil
 * @create 2019/3/16
 */
func login(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/login" //填空没有默认值
	params := map[string]string{
		"user[李长全]": "123456",
	}
	send := postSend(url, params)
	fmt.Println(send)
}

func get(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/show?lastname=nht&pri[2]=3&list=[1,2,3,4]" //填空没有默认值
	result := queryGet(t, url)
	fmt.Println(result)
}

func somePost(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/somePost"
	params := map[string]string{
		"message": "222",
		"nick":    "", //这种形式也算有值,不会填充默认值
		"pri[2]":  "2",
		"list":    `[1,2,3,4]`,
	}

	send := postSend(url, params)
	fmt.Println(send)
}

func put(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/somePut" //填空没有默认值
	client := &http.Client{}
	result := puts(t, client, url)
	fmt.Println(result)
}

func patch(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/somePatch" //填空没有默认值
	client := &http.Client{}
	result := patchs(t, client, url)
	fmt.Println(result)
}

func delete(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/someDelete" //填空没有默认值
	client := &http.Client{}
	result := deletes(t, client, url)
	fmt.Println(result)
}

func head(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/someHead" //填空没有默认值
	client := &http.Client{}
	heads(t, client, url)
}

func someOptions(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/pencil/someOptions" //填空没有默认值
	client := &http.Client{}
	result := option(t, client, url)
	fmt.Println(result)
}
