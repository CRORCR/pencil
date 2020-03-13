package test

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

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

func postSendCopy(url string, params interface{}) string {
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

func postSendList(url string, params map[string][]string) string {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range params {
		for _, val := range value {
			_ = writer.WriteField(key, val)
		}
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
