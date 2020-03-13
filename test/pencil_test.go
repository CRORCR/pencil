package test

import (
	"testing"
)

/**
 * @desc  测试
 * @author Ipencil
 * @create 2019/3/15
 */
func TestGet(t *testing.T) {
	t.Run("reqType", reqType)
	t.Run("upload", upload)
	t.Run("upload_more", upload_more)
	t.Run("band", band)
	t.Run("filter", filter)
}
