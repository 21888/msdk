// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description: 键盘操作函数
package msdk

import (
	"syscall"
	"unsafe"
)

var (
	//int WINAPI M_KeyPress(HANDLE m_hdl, int HidKeyCode, int Nbr);
	mKeyPress = deviceDll.NewProc("M_KeyPress")
	//int WINAPI M_KeyDown(HANDLE m_hdl, int HidKeyCode);
	mKeyDown = deviceDll.NewProc("M_KeyDown")
	//int WINAPI M_KeyUp(HANDLE m_hdl, int HidKeyCode);
	mKeyUp = deviceDll.NewProc("M_KeyUp")
	//int WINAPI M_KeyState(HANDLE m_hdl, int HidKeyCode);
	mKeyState = deviceDll.NewProc("M_KeyState")
	//int WINAPI M_KeyPress2(HANDLE m_hdl, int KeyCode, int Nbr);
	mKeyPress2 = deviceDll.NewProc("M_KeyPress2")
	//int WINAPI M_KeyDown2(HANDLE m_hdl, int KeyCode);
	mKeyDown2 = deviceDll.NewProc("M_KeyDown2")
	//int WINAPI M_KeyUp2(HANDLE m_hdl, int KeyCode);
	mKeyUp2 = deviceDll.NewProc("M_KeyUp2")
	//int WINAPI M_KeyState2(HANDLE m_hdl, int KeyCode);
	mKeyState2 = deviceDll.NewProc("M_KeyState2")
	//int WINAPI M_ReleaseAllKey(HANDLE m_hdl);
	mReleaseAllKey = deviceDll.NewProc("M_ReleaseAllKey")
	//int WINAPI M_NumLockLedState(HANDLE m_hdl);
	mNumLockLedState = deviceDll.NewProc("M_NumLockLedState")
	//int WINAPI M_CapsLockLedState(HANDLE m_hdl);
	mCapsLockLedState = deviceDll.NewProc("M_CapsLockLedState")
	//int WINAPI M_ScrollLockLedState(HANDLE m_hdl);
	mScrollLockLedState = deviceDll.NewProc("M_ScrollLockLedState")
	//int WINAPI M_KeyInputString(HANDLE m_hdl, char *InputStr, int InputLen);
	mKeyInputString = deviceDll.NewProc("M_KeyInputString")
	//int WINAPI M_KeyInputStringGBK(HANDLE m_hdl, char *InputStr, int InputLen);
	mKeyInputStringGBK = deviceDll.NewProc("M_KeyInputStringGBK")
	//int WINAPI M_KeyInputStringUnicode(HANDLE m_hdl, char *InputStr, int InputLen);
	mKeyInputStringUnicode = deviceDll.NewProc("M_KeyInputStringUnicode")
)

///////////以下接口采用的HidKeyCode是USB键盘的KeyCode，不是windows的KeyCode，新开发程序不建议再使用

// MKeyPress 单击(按下后立刻弹起)按键
// 单击(按下后立刻弹起)按键  //HidKeyCode: 键盘码; Nbr: 按下次数
func MKeyPress(hdl HANDLE, hidKeyCode, nbr int) (int, error) {
	ret, err := callProc(mKeyPress, uintptr(hdl), uintptr(hidKeyCode), uintptr(nbr))
	return int(ret), err
}

// MKeyDown 按下按键
// 按下按键  //HidKeyCode: 键盘码
func MKeyDown(hdl HANDLE, hidKeyCode int) (int, error) {
	ret, err := callProc(mKeyDown, uintptr(hdl), uintptr(hidKeyCode))
	return int(ret), err
}

// MKeyUp 弹起按键
// 弹起按键  //HidKeyCode: 键盘码
func MKeyUp(hdl HANDLE, hidKeyCode int) (int, error) {
	ret, err := callProc(mKeyUp, uintptr(hdl), uintptr(hidKeyCode))
	return int(ret), err
}

// MKeyState 按键状态
// 按键状态  //HidKeyCode: 键盘码
func MKeyState(hdl HANDLE, hidKeyCode int) (int, error) {
	ret, err := callProc(mKeyState, uintptr(hdl), uintptr(hidKeyCode))
	return int(ret), err
}

// MKeyPress2 单击(按下后立刻弹起)按键
// 单击(按下后立刻弹起)按键  //KeyCode: 键盘码; Nbr: 按下次数
func MKeyPress2(hdl HANDLE, keyCode, nbr int) (int, error) {
	ret, err := callProc(mKeyPress2, uintptr(hdl), uintptr(keyCode), uintptr(nbr))
	return int(ret), err
}

// MKeyDown2 按下按键
// 按下按键  //KeyCode: 键盘码
func MKeyDown2(hdl HANDLE, keyCode int) (int, error) {
	ret, err := callProc(mKeyDown2, uintptr(hdl), uintptr(keyCode))
	return int(ret), err
}

// MKeyUp2 弹起按键
// 弹起按键  //KeyCode: 键盘码
func MKeyUp2(hdl HANDLE, keyCode int) (int, error) {
	ret, err := callProc(mKeyUp2, uintptr(hdl), uintptr(keyCode))
	return int(ret), err
}

// MKeyState2 按键状态
// 按键状态  //KeyCode: 键盘码
func MKeyState2(hdl HANDLE, keyCode int) (int, error) {
	ret, err := callProc(mKeyState2, uintptr(hdl), uintptr(keyCode))
	return int(ret), err
}

// MReleaseAllKey 释放所有按键
// 释放所有按键
func MReleaseAllKey(hdl HANDLE) (int, error) {
	ret, err := callProc(mReleaseAllKey, uintptr(hdl))
	return int(ret), err
}

// MNumLockLedState 数字键灯状态
// 数字键灯状态
func MNumLockLedState(hdl HANDLE) (int, error) {
	ret, err := callProc(mNumLockLedState, uintptr(hdl))
	return int(ret), err
}

// MCapsLockLedState 大小写灯状态
// 大小写灯状态
func MCapsLockLedState(hdl HANDLE) (int, error) {
	ret, err := callProc(mCapsLockLedState, uintptr(hdl))
	return int(ret), err
}

// MScrollLockLedState 滚动键灯状态
// 滚动键灯状态
func MScrollLockLedState(hdl HANDLE) (int, error) {
	ret, err := callProc(mScrollLockLedState, uintptr(hdl))
	return int(ret), err
}

// MKeyInputString 键盘输入字符串
// 键盘输入字符串  //InputStr: 输入字符串; InputLen: 字符串长度
func MKeyInputString(hdl HANDLE, inputStr string, inputLen int) (int, error) {
	ret, err := callProc(mKeyInputString, uintptr(hdl), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(inputStr))), uintptr(inputLen))
	return int(ret), err
}

// MKeyInputStringGBK 键盘输入字符串
// 键盘输入字符串  //InputStr: 输入字符串; InputLen: 字符串长度
func MKeyInputStringGBK(hdl HANDLE, inputStr string, inputLen int) (int, error) {
	ret, err := callProc(mKeyInputStringGBK, uintptr(hdl), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(inputStr))), uintptr(inputLen))
	return int(ret), err
}

// MKeyInputStringUnicode 键盘输入字符串
// 键盘输入字符串  //InputStr: 输入字符串; InputLen: 字符串长度
func MKeyInputStringUnicode(hdl HANDLE, inputStr string, inputLen int) (int, error) {
	ret, err := callProc(mKeyInputStringUnicode, uintptr(hdl), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(inputStr))), uintptr(inputLen))
	return int(ret), err
}
