// Package msdk author: 茶白
// CreateDate: 2023/5/6
// Description: 鼠标操作
package msdk

var (
	//int WINAPI M_LeftClick(HANDLE m_hdl,int Nbr);
	mLeftClick = deviceDll.NewProc("M_LeftClick")
	//int WINAPI M_LeftDoubleClick(HANDLE m_hdl,int Nbr);
	mLeftDoubleClick = deviceDll.NewProc("M_LeftDoubleClick")
	//int WINAPI M_LeftDown(HANDLE m_hdl);
	mLeftDown = deviceDll.NewProc("M_LeftDown")
	//int WINAPI M_LeftUp(HANDLE m_hdl);
	mLeftUp = deviceDll.NewProc("M_LeftUp")
	//int WINAPI M_RightClick(HANDLE m_hdl,int Nbr);
	mRightClick = deviceDll.NewProc("M_RightClick")
	//int WINAPI M_RightDown(HANDLE m_hdl);
	mRightDown = deviceDll.NewProc("M_RightDown")
	//int WINAPI M_RightUp(HANDLE m_hdl);
	mRightUp = deviceDll.NewProc("M_RightUp")
	//int WINAPI M_MiddleClick(HANDLE m_hdl,int Nbr);
	mMiddleClick = deviceDll.NewProc("M_MiddleClick")
	//int WINAPI M_MiddleDown(HANDLE m_hdl);
	mMiddleDown = deviceDll.NewProc("M_MiddleDown")
	//int WINAPI M_MiddleUp(HANDLE m_hdl);
	mMiddleUp = deviceDll.NewProc("M_MiddleUp")
	//int WINAPI M_ReleaseAllMouse(HANDLE m_hdl);
	mReleaseAllMouse = deviceDll.NewProc("M_ReleaseAllMouse")
	//int WINAPI M_MouseKeyState(HANDLE m_hdl, int MouseKeyCode);
	mMouseKeyState = deviceDll.NewProc("M_MouseKeyState")
	//int WINAPI M_MouseWheel(HANDLE m_hdl,int Nbr);
	mMouseWheel = deviceDll.NewProc("M_MouseWheel")
	//int WINAPI M_ResetMousePos(HANDLE m_hdl);
	mResetMousePos = deviceDll.NewProc("M_ResetMousePos")
	//int WINAPI M_MoveR(HANDLE m_hdl,int x, int y);
	mMoveR = deviceDll.NewProc("M_MoveR")
	//int WINAPI M_MoveTo(HANDLE m_hdl,int x, int y);
	mMoveTo = deviceDll.NewProc("M_MoveTo")
	//int WINAPI M_GetCurrMousePos(HANDLE m_hdl,int *x, int *y);
	mGetCurrMousePos = deviceDll.NewProc("M_GetCurrMousePos")
	//int WINAPI M_GetCurrMousePosX(HANDLE m_hdl); //由于某些语言不支持指针，增加此接口
	mGetCurrMousePosX = deviceDll.NewProc("M_GetCurrMousePosX")
	//int WINAPI M_GetCurrMousePosY(HANDLE m_hdl); //由于某些语言不支持指针，增加此接口
	mGetCurrMousePosY = deviceDll.NewProc("M_GetCurrMousePosY")
	//int WINAPI M_MoveR2(HANDLE m_hdl,int x, int y);
	mMoveR2 = deviceDll.NewProc("M_MoveR2")
	//int WINAPI M_MoveTo2(HANDLE m_hdl,int x, int y);
	mMoveTo2 = deviceDll.NewProc("M_MoveTo2")
	//int WINAPI M_GetCurrMousePos2(int *x, int *y);
	mGetCurrMousePos2 = deviceDll.NewProc("M_GetCurrMousePos2")
	//int WINAPI M_ResolutionUsed(HANDLE m_hdl, int x, int y);
	mResolutionUsed = deviceDll.NewProc("M_ResolutionUsed")
	//int WINAPI M_MoveTo3(HANDLE m_hdl, int x, int y);
	mMoveTo3 = deviceDll.NewProc("M_MoveTo3")
	//int WINAPI M_MoveTo3_D(HANDLE m_hdl, int x, int y);
	mMoveTo3_D = deviceDll.NewProc("M_MoveTo3_D")
)

//	 MLeftClick 左键单击
//		左键在当前位置单击次数
func MLeftClick(m_hdl HANDLE, nbr int) (int, error) {
	ret, err := callProc(mLeftClick, uintptr(m_hdl), uintptr(nbr))
	return int(ret), err
}

//	 MLeftDoubleClick 左键双击
//		左键在当前位置双击次数
func MLeftDoubleClick(m_hdl HANDLE, nbr int) (int, error) {
	ret, err := callProc(mLeftDoubleClick, uintptr(m_hdl), uintptr(nbr))
	return int(ret), err
}

// MLeftDown 按下左键不弹起
func MLeftDown(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mLeftDown, uintptr(m_hdl))
	return int(ret), err
}

// MLeftUp 弹起左键
func MLeftUp(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mLeftUp, uintptr(m_hdl))
	return int(ret), err
}

//	 MRightClick 右键单击
//		左键在当前位置单击次数
func MRightClick(m_hdl HANDLE, nbr int) (int, error) {
	ret, err := callProc(mRightClick, uintptr(m_hdl), uintptr(nbr))
	return int(ret), err
}

// MRightDown 按下右键不弹起
func MRightDown(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mRightDown, uintptr(m_hdl))
	return int(ret), err
}

// MRightUp 弹起右键
func MRightUp(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mRightUp, uintptr(m_hdl))
	return int(ret), err
}

//	 MMiddleClick 中键单击
//		左键在当前位置单击次数
func MMiddleClick(m_hdl HANDLE, nbr int) (int, error) {
	ret, err := callProc(mMiddleClick, uintptr(m_hdl), uintptr(nbr))
	return int(ret), err
}

// MMiddleDown 按下中键不弹起
func MMiddleDown(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mMiddleDown, uintptr(m_hdl))
	return int(ret), err
}

// MMiddleUp 弹起中键
func MMiddleUp(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mMiddleUp, uintptr(m_hdl))
	return int(ret), err
}

// MReleaseAllMouse 弹起鼠标的所有按键
func MReleaseAllMouse(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mReleaseAllMouse, uintptr(m_hdl))
	return int(ret), err
}

//	 MMouseKeyState 读取鼠标左中右键状态
//		MouseKeyCode: 1=左键 2=右键 3=中键
//		返回 0: 弹起状态；1:按下状态；-1: 失败
//		只能读取盒子中鼠标的状态，读取不到实体鼠标的状态
func MMouseKeyState(m_hdl HANDLE, MouseKeyCode int) (int, error) {
	ret, err := callProc(mMouseKeyState, uintptr(m_hdl), uintptr(MouseKeyCode))
	return int(ret), err
}

//	 MMouseWheel 滚动鼠标滚轮
//		Nbr: 滚动量,  为正,向上滚动；为负, 向下滚动;
func MMouseWheel(m_hdl HANDLE, Nbr int) (int, error) {
	ret, err := callProc(mMouseWheel, uintptr(m_hdl), uintptr(Nbr))
	return int(ret), err
}

//	 MResetMousePos 将鼠标移动到原点(0,0)
//		在出现移动出现异常时，可以用该函数将鼠标复位
func MResetMousePos(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mResetMousePos, uintptr(m_hdl))
	return int(ret), err
}

//	 MMoveR 相对移动鼠标
//		从当前位置移动鼠标    x: x方向（横轴）的距离（正:向右; 负值:向左）; y: y方向（纵轴）的距离（正:向下; 负值:向上）
func MMoveR(m_hdl HANDLE, x, y int) (int, error) {
	ret, err := callProc(mMoveR, uintptr(m_hdl), uintptr(x), uintptr(y))
	return int(ret), err
}

//	 MMoveTo 移动鼠标到指定坐标
//		移动鼠标到指定坐标     x: x方向（横轴）的坐标; y: y方向（纵轴）的坐标。坐标原点(0, 0)在屏幕左上角
//		注意：如果出现过将鼠标移动的距离超过屏幕大小，再次MoveTo可能会出现无法正确移动到指定坐标的问题，如果出现该问题，需调用ResetMousePos复位
func MMoveTo(m_hdl HANDLE, x, y int) (int, error) {
	ret, err := callProc(mMoveTo, uintptr(m_hdl), uintptr(x), uintptr(y))
	return int(ret), err
}

//	  MGetCurrMousePos 读取当前鼠标所在坐标  返回坐标在x、y中。
//		注意：该函数必须在执行一次MoveTo或ResetMousePos函数后才能正确执行！
//		注意：如果曾经出现过将鼠标移动的距离超过屏幕大小，这里读取到的坐标值有可能是不正确的！如果出现该问题，需调用ResetMousePos复位
func MGetCurrMousePos(m_hdl HANDLE) (x, y int, err error) {
	ret, err := callProc(mGetCurrMousePos, uintptr(m_hdl))
	return int(ret >> 16), int(ret & 0xffff), err
}

// MGetCurrMousePosX 返回当前鼠标坐标X值
func MGetCurrMousePosX(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mGetCurrMousePosX, uintptr(m_hdl))
	return int(ret), err
}

// MGetCurrMousePosY 返回当前鼠标坐标Y值
func MGetCurrMousePosY(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mGetCurrMousePosY, uintptr(m_hdl))
	return int(ret), err
}

////////以下接口仅适用主控机和被控机是同一台电脑的使用方式(单头模块；双头模块的两个USB头都连接到同一台电脑)
////////以下接口将调用系统的API来获取当前鼠标位置，DLL将不记录鼠标移动的位置

//	 MMoveR2 移动鼠标到指定坐标
//		移动鼠标到指定坐标    x: x方向（横轴）的坐标; y: y方向（纵轴）的坐标。
func MMoveR2(x, y int) (int, error) {
	ret, err := callProc(mMoveR2, uintptr(x), uintptr(y))
	return int(ret), err
}

//	 MMoveTo2 移动鼠标到指定坐标
//		移动鼠标到指定坐标    x: x方向（横轴）的坐标; y: y方向（纵轴）的坐标。坐标原点(0, 0)在屏幕左上角
func MMoveTo2(x, y int) (int, error) {
	ret, err := callProc(mMoveTo2, uintptr(x), uintptr(y))
	return int(ret), err
}

// MGetCurrMousePos2 读取当前鼠标所在坐标  返回坐标在x、y中。
func MGetCurrMousePos2() (x, y int, err error) {
	ret, err := callProc(mGetCurrMousePos2)
	return int(ret >> 16), int(ret & 0xffff), err
}

////////以下接口将使用绝对移动功能。该功能目前还不能支持安卓

//	 MResolutionUsed 输入被控机分辨率
//		在使用绝对移动功能前，必须先输入被控机的屏幕分辨率，打开端口获取句柄后，只需要调用一次该接口就可以
//		x: x方向（横轴）的坐标; y: y方向（纵轴）的坐标。坐标原点(0, 0)在屏幕左上角
//		返回值如果是-10，表示该盒子不支持绝对移动功能。返回0表示执行正确。可以用该接口判断盒子是否支持绝对移动功能
func MResolutionUsed(m_hdl HANDLE) (int, error) {
	ret, err := callProc(mResolutionUsed, uintptr(m_hdl))
	return int(ret), err
}

//	 MMoveTo3 将鼠标移动到指定坐标。
//		  绝对移动功能，鼠标移动到指定位置时，在某些坐标上最大会有±2的误差
//		  使用该接口后，可以调用M_GetCurrMousePos读取鼠标位置
func MMoveTo3(x, y int) (int, error) {
	ret, err := callProc(mMoveTo3, uintptr(x), uintptr(y))
	return int(ret), err
}

//		MMoveTo3D 一步到位将鼠标移动到指定坐标。
//	 	   使用绝对移动功能，鼠标一步到位移动（没有移动轨迹）到指定位置，在某些坐标上最大会有±2的误差
//		   使用该接口后，可以调用M_GetCurrMousePos读取鼠标位置
func MMoveTo3D(x, y int) (int, error) {
	ret, err := callProc(mMoveTo3_D, uintptr(x), uintptr(y))
	return int(ret), err
}
