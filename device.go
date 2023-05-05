// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description: 设备操作函数
package msdk

import (
	"syscall"
	"unsafe"
)

var (
	mOpen        = deviceDll.NewProc("M_Open")
	mOpenVidPid  = deviceDll.NewProc("M_Open_VidPid")
	mScanAndOpen = deviceDll.NewProc("M_ScanAndOpen")
	mClose       = deviceDll.NewProc("M_Close")

	//int WINAPI M_GetDevSn(HANDLE m_hdl, DWORD *dwp_LenResponse, unsigned char *ucp_Response);
	mGetDevSn = deviceDll.NewProc("M_GetDevSn")
)

//	 callProc 调用 LazyProc.Call，检查返回值并返回错误。
//		中间件，用于检查返回值
func callProc(proc *syscall.LazyProc, args ...uintptr) (uintptr, error) {
	ret, _, _ := proc.Call(args...)
	/*if err != nil {
		return ret, syscall.Errno(ret)
	}*/
	//fmt.Println("callProc", ret, err)
	return ret, nil
}

//	 MOpen 打开端口获取句柄;
//		Nbr是端口号，无论是双头模块还是单头模块，都是从1开始，依次为2/3/4...，最大126。电脑上插入一个就始终是1；插入n个 端口分别是1、2....n
//		使用M_Open接口打开默认vid pid的盒子
func MOpen(Nbr int) (HANDLE, error) {
	ret, err := callProc(mOpen, uintptr(Nbr))
	return HANDLE(ret), err
}

//	 MOpenVidpid 打开指定vid pid的盒子
//		使用M_Open_VidPid接口可以打开指定vid pid的单头盒子或者双头盒子的主控端。必须保证电脑上只插有一个这种盒子。
func MOpenVidpid(Vid, Pid int) (HANDLE, error) {
	ret, err := callProc(mOpenVidPid, uintptr(Vid), uintptr(Pid))
	return HANDLE(ret), err
}

//	 MScanAndOpen
//		扫描硬件列表，打开扫描到的第一个单头或者双头盒子
func MScanAndOpen() (HANDLE, error) {
	ret, err := callProc(mScanAndOpen)
	return HANDLE(ret), err
}

//	 MClose
//		关闭端口；在程序退出前再关闭端口; 返回 0: 成功；!0: 失败
func MClose(hdl HANDLE) error {
	_, err := callProc(mClose, uintptr(hdl))
	return err
}

//	 MGetDevSn
//		获取设备序列号
//		参数:
//			dwp_LenResponse: 设备序列号的长度(单位: 字节)
//			ucp_Response: 设备序列号buf(buf由调用该API的应用程序分配)
//		返回 0: 成功；-1: 失败
func MGetDevSn(hdl HANDLE, dwp_LenResponse *uint32, ucp_Response *byte) error {
	_, err := callProc(mGetDevSn, uintptr(hdl), uintptr(unsafe.Pointer(dwp_LenResponse)), uintptr(unsafe.Pointer(ucp_Response)))
	return err
}
