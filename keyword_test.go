// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description:
package msdk

import "testing"

func TestMKeyInputStringUnicode(t *testing.T) {
	if pid, err := MOpen(1); err != nil {
		t.Error(err)
	} else {
		t.Log(pid)
		str := "Hello world!"
		MKeyInputStringUnicode(pid, str, len(str))
		_ = MClose(pid) // 关闭端口
	}
}
