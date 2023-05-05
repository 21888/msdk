// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description:
package msdk

import (
	"testing"
)

func TestM_Open(t *testing.T) {
	if pid, err := MOpen(1); err != nil {
		t.Error(err)
	} else {
		t.Log(pid)
	}
}
func TestM_Open_VidPid(t *testing.T) {
	if pid, err := MOpenVidpid(0xC216, 0x0301); err != nil {
		t.Error(err)
	} else {
		t.Log("句柄", pid)
		_ = MClose(pid) // 关闭端口
	}
}
func TestMGetDevSn(t *testing.T) {
	if pid, err := MOpenVidpid(0xC216, 0x0301); err != nil {
		t.Error(err)
	} else {
		t.Log(pid)
		var xLe uint32
		const bufferLen = 256 // 替换为所需的实际缓冲区长度
		buf := make([]byte, bufferLen)
		MGetDevSn(pid, &xLe, &buf[0])
		t.Log("序列号: ", xLe, buf)
		_ = MClose(pid) // 关闭端口
	}
}
