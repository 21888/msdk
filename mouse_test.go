// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description:
package msdk

import (
	"testing"
)

func TestMMoveR(t *testing.T) {
	if pid, err := MOpen(1); err != nil {
		t.Error(err)
	} else {
		t.Log(pid)
		MMoveR(pid, 100, 100)
		_ = MClose(pid) // 关闭端口
	}
}
