package show

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
 * @desc 单文件上传 \ 多文件上传  图片
 * @author Ipencil
 * @create 2019/3/15
 */
func UploadOne(c *gin.Context) {
	file, err := c.FormFile("files")
	if err != nil {
		fmt.Println("读取文件失败")
	}
	right := strings.Replace(file.Filename, "\\", " ", -1)
	fmt.Println("文件名称:", right[len(right)-1])
	fmt.Println("文件大小:", file.Size)
	file2, err := os.OpenFile("hello/test/aa.jpg", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("创建文件失败")
	}
	open, err := file.Open()
	if err != nil {
		fmt.Println("open failed")
	}
	defer func() {
		open.Close()
		file2.Close()
	}()
	io.Copy(file2, open)
	c.JSON(http.StatusOK, gin.H{"result": "上传图片成功", "error_code": "success"})
	return
}

func UploadAll(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for i := 0; i < len(files); i++ {
		right := strings.Replace(files[i].Filename, "\\", " ", -1)
		fmt.Println("文件名称:", right[len(right)-1])
		rand.Intn(10)
		file2, err := os.OpenFile(fmt.Sprintf("hello/test/aa%v.jpg", i), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("创建文件失败")
		}
		open, err := files[i].Open()
		if err != nil {
			fmt.Println("open failed")
		}
		defer func() {
			open.Close()
			file2.Close()
		}()
		io.Copy(file2, open)
	}

	c.JSON(http.StatusOK, gin.H{"result": "上传图片成功", "error_code": "success"})
	return
}
