// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description: 通用操作函数
package msdk

import "syscall"

// MSuccess 表示调用成功的返回值。
const MSuccess = 0

// HANDLE 表示 mdevice.dll 中返回的句柄类型。
type HANDLE uintptr

var (
	deviceDll = syscall.NewLazyDLL("msdk64.dll")

	//int WINAPI M_Delay(int time);
	mDelay = deviceDll.NewProc("M_Delay")
	//int WINAPI M_DelayRandom(int Min_time, int Max_time);
	mDelayRandom = deviceDll.NewProc("M_DelayRandom")
	//int WINAPI M_RandDomNbr(int Min_V, int Max_V);
	mRandDomNbr = deviceDll.NewProc("M_RandDomNbr")
)

// MDelay 延迟
// 延时指定时间  time:单位ms
func MDelay(time int) (int, error) {
	ret, err := callProc(mDelay, uintptr(time))
	return int(ret), err
}

// MDelayRandom 随机延迟
// 在指定的最小最大值之间延时随机时间  Min_time:最小延时时间; Max_time: 最大延时时间 （单位：ms）
func MDelayRandom(minTime, maxTime int) (int, error) {
	ret, err := callProc(mDelayRandom, uintptr(minTime), uintptr(maxTime))
	return int(ret), err
}

// MRandDomNbr 随机数
// 在最小最大值之间取随机数
func MRandDomNbr(minV, maxV int) (int, error) {
	ret, err := callProc(mRandDomNbr, uintptr(minV), uintptr(maxV))
	return int(ret), err
}
