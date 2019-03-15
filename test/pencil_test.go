package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

/*
get或者post可以发送map结构和数组

 */
const LICHANGQUAN = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MCwiTmFtZSI6IuadjumVv-WFqCIsIlBhc3N3b3JkIjoiIiwiU2hhcmUiOiIxMjM0NTYiLCJJY29uIjoiIiwiZXhwIjoxNTYwNDE2MDU2LCJpc3MiOiJwZW5jaWwiLCJuYmYiOjE1NTI2NDAwNTZ9.uBqblo8ENAsf3yNyCUPw2oPIK5Pt98GPfkPp2ewgjJs`

/**
 * @desc  测试
 * @author Ipencil
 * @create 2019/3/15
 */
func TestGet(t *testing.T) {
	t.Run("login", login)
	t.Run("somePost", somePost)
	t.Run("get", get)
	t.Run("band_json", band_json)
	t.Run("band_xml", band_xml)

	t.Run("put", put)
	t.Run("patch", patch)
	t.Run("delete", delete)
	t.Run("head", head)
	t.Run("someOptions", someOptions)
	t.Run("upload", upload)
	t.Run("uploada", uploada)
}

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

func band_json(t *testing.T) {
	//t.SkipNow()
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

func band_xml(t *testing.T) {
	t.SkipNow()
	url := "http://localhost:8080/bind"
	params := map[string]string{
		"name":     "lcq",
		"password": "1234", //这种形式也算有值,不会填充默认值
	}
	send := postSend(url, params)
	fmt.Println(send)
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

func puts(t *testing.T, client *http.Client, url string) string {
	request, err := http.NewRequest("PUT", url, nil)
	request.Header.Set("Authorization", LICHANGQUAN)
	assertNil(t, "", err)
	//处理返回
	response, _ := client.Do(request)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
	return string(bytes)
}

func patchs(t *testing.T, client *http.Client, url string) string {
	request, err := http.NewRequest("PATCH", url, nil)
	request.Header.Set("Authorization", LICHANGQUAN)
	assertNil(t, "", err)
	//处理返回
	response, _ := client.Do(request)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
	return string(bytes)
}

func deletes(t *testing.T, client *http.Client, url string) string {
	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Authorization", LICHANGQUAN)
	assertNil(t, "", err)
	//处理返回
	response, _ := client.Do(request)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
	return string(bytes)
}

func heads(t *testing.T, client *http.Client, url string) {
	request, err := http.NewRequest("HEAD", url, nil)
	request.Header.Set("Authorization", LICHANGQUAN)
	assertNil(t, "", err)
	//处理返回
	response, _ := client.Do(request)
	defer response.Body.Close()
	//解析list
	t.Logf("result:%+v\n", response.Header)
}

// "", "PROPFIND", "SEARCH"
func option(t *testing.T, client *http.Client, url string) string {
	request, err := http.NewRequest("OPTIONS", url, nil)
	request.Header.Set("Authorization", LICHANGQUAN)
	assertNil(t, "", err)
	//处理返回
	response, _ := client.Do(request)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
	return string(bytes)
}

func postSend(url string,params map[string]string,)string{
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()

	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", LICHANGQUAN)

	response, _ := client.Do(request)
	defer func() { response.Body.Close() }()
	bytes, _ := ioutil.ReadAll(response.Body)
	return string(bytes)
}

func queryGet(t *testing.T, url string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", LICHANGQUAN)
	assertNil(t, "", err)
	//处理返回
	response, _ := client.Do(request)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
	return string(bytes)
}

func assertNil(t *testing.T, name string, v ...interface{}) {
	for _, value := range v {
		if value == nil {
			continue
		}
		if name == "" {
			t.Errorf("Not Nil %v", value)
		} else {
			t.Errorf("%s %v", name, value)
		}
	}
}
