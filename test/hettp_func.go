package test

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"pencil/api/bind"
	"testing"
)

const LICHANGQUAN = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MCwiTmFtZSI6IuadjumVv-WFqCIsIlBhc3N3b3JkIjoiIiwiU2hhcmUiOiIxMjM0NTYiLCJJY29uIjoiIiwiZXhwIjoxNTYwNDE2MDU2LCJpc3MiOiJwZW5jaWwiLCJuYmYiOjE1NTI2NDAwNTZ9.uBqblo8ENAsf3yNyCUPw2oPIK5Pt98GPfkPp2ewgjJs`

/**
 * @desc    所有的请求类型 测试
 * @author Ipencil
 * @create 2019/3/16
 */

func reqType(t *testing.T) {
	t.Run("get", get)
	t.Run("put", put)
	t.Run("patch", patch)
	t.Run("delete", delete)
	t.Run("head", head)
	t.Run("someOptions", someOptions)
	t.Run("somePost", somePost)
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

func postSendCopy(url string, params bind.User) string {
	client := &http.Client{}
	body := &bytes.Buffer{}
	byt, err := xml.Marshal(params)
	if err != nil {
		return ""
	}
	body.Write(byt)
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Authorization", LICHANGQUAN)
	request.Header.Set("Content-Type", "application/xml")
	response, _ := client.Do(request)
	defer func() { response.Body.Close() }()
	bytes, _ := ioutil.ReadAll(response.Body)
	return string(bytes)
}

func postSend(url string, params map[string]string) string {
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
