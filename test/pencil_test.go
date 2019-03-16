package test

import (
	"testing"
)

/*
get或者post可以发送map结构和数组

*/

/**
 * @desc  测试
 * @author Ipencil
 * @create 2019/3/15
 */
func TestGet(t *testing.T) {
	t.Run("login", login)
	t.Run("band", band)
	t.Run("reqType", reqType)
	t.Run("upload", upload)
	t.Run("upload_more", upload_more)
}
