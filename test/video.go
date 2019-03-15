package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/15
 */
func upload(t *testing.T) {
	t.SkipNow()
	client := &http.Client{}
	params := map[string]string{
	}

	upurl := "http://localhost:8080/pencil/upload"
	request, e := newfileUploadRequest(upurl, params)
	if e != nil {
		fmt.Println("error", e)
		return
	}
	response, _ := client.Do(request)
	defer func() { response.Body.Close() }()
	assert.Equal(t, "200 OK", response.Status)
	bytes, _ := ioutil.ReadAll(response.Body)
	t.Log("result:", string(bytes))
}

//多图片上传
func uploada(t *testing.T) {
	t.SkipNow()
	client := &http.Client{}
	params := map[string]string{
	}

	upurl := "http://localhost:8080/pencil/uploada"
	request, e := newfileUploadRequest(upurl, params)
	if e != nil {
		fmt.Println("error", e)
		return
	}
	response, _ := client.Do(request)
	defer func() { response.Body.Close() }()
	assert.Equal(t, "200 OK", response.Status)
	bytes, _ := ioutil.ReadAll(response.Body)
	t.Log("result:", string(bytes))
}

func newfileUploadRequest(uri string, params map[string]string) (*http.Request, error) {
	filePaths := []string{
		//"K:\\upload\\店铺\\店铺0.jpg",
		//"K:\\upload\\店铺\\店铺1.jpg",
		//"K:\\upload\\店铺\\店铺2.jpg",
		//"K:\\upload\\轮播\\轮播0.jpg",
		"K:\\upload\\轮播\\轮播1.jpg",
		"K:\\upload\\轮播\\轮播2.jpg",
		//"K:\\upload\\视频\\bb.mp4",
	}
	keys := []string{
		"files",
		//"infoimg1",
		//"infoimg2",
		//"loopimg0",
		//"loopimg1",
		//"loopimg2",
		//"video0",
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var part io.Writer
	for i := 0; i < len(filePaths); i++ {
		file, err := os.Open(filePaths[i])
		if err != nil {
			return nil, err
		}
		if strings.Contains(uri, "uploada") {
			part, err = writer.CreateFormFile("upload[]", filePaths[i])
		} else {
			part, err = writer.CreateFormFile(keys[i], filePaths[i])
		}
		_, err = io.Copy(part, file)
		file.Close()
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, err
}
