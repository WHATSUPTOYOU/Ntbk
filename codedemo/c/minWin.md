```C

// Coverup.cpp : 定义应用程序的入口点。
//

#include "framework.h"
#include "Coverup.h"
#include<string>

#define MAX_LOADSTRING 100

// 全局变量:
HINSTANCE hInst;                                // 当前实例
WCHAR szTitle[MAX_LOADSTRING];                  // 标题栏文本
WCHAR szWindowClass[MAX_LOADSTRING];            // 主窗口类名

// 此代码模块中包含的函数的前向声明:
ATOM                MyRegisterClass(HINSTANCE hInstance);
BOOL                InitInstance(HINSTANCE, int);
LRESULT CALLBACK    WndProc(HWND, UINT, WPARAM, LPARAM);

int APIENTRY wWinMain(_In_ HINSTANCE hInstance,
                     _In_opt_ HINSTANCE hPrevInstance,
                     _In_ LPWSTR    lpCmdLine,
                     _In_ int       nCmdShow)
{
	//最小化桌面窗口,直到最上层为任务栏
	char szTitleWin[MAX_PATH] = { 0 };
	char szClass[MAX_PATH] = { 0 };
	int nMaxCount = MAX_PATH;

	LPSTR lpClassName = szClass;
	LPSTR lpWindowName = szTitleWin;

	while (true) {
		HWND hwnd = GetForegroundWindow();//最前端的窗口 

		GetWindowTextA(hwnd, lpWindowName, nMaxCount);
		GetClassNameA(hwnd, lpClassName, nMaxCount);
		std::string str = (std::string)lpClassName;

		if (str.find("TrayWnd") != std::string::npos) {
			break;
		}

		PostMessage(hwnd, WM_SYSCOMMAND, SC_MINIMIZE, 0);
		Sleep(100);
	}

    UNREFERENCED_PARAMETER(hPrevInstance);
    UNREFERENCED_PARAMETER(lpCmdLine);

    // TODO: 在此处放置代码。

    // 初始化全局字符串
    LoadStringW(hInstance, IDS_APP_TITLE, szTitle, MAX_LOADSTRING);
    LoadStringW(hInstance, IDC_COVERUP, szWindowClass, MAX_LOADSTRING);
    MyRegisterClass(hInstance);

    // 执行应用程序初始化:
    if (!InitInstance (hInstance, SW_MAXIMIZE))
    {
        return FALSE;
    }

    HACCEL hAccelTable = LoadAccelerators(hInstance, MAKEINTRESOURCE(IDC_COVERUP));

    MSG msg;

     主消息循环:
    while (GetMessage(&msg, nullptr, 0, 0))
    {
        if (!TranslateAccelerator(msg.hwnd, hAccelTable, &msg))
        {
            TranslateMessage(&msg);
            DispatchMessage(&msg);
        }
    }

    return 0;
}



//
//  函数: MyRegisterClass()
//
//  目标: 注册窗口类。
//
ATOM MyRegisterClass(HINSTANCE hInstance)
{
    WNDCLASSEXW wcex;

    wcex.cbSize = sizeof(WNDCLASSEX);

    wcex.style          = CS_HREDRAW | CS_VREDRAW;
    wcex.lpfnWndProc    = WndProc;
    wcex.cbClsExtra     = 0;
    wcex.cbWndExtra     = 0;
    wcex.hInstance      = hInstance;
    wcex.hIcon          = LoadIcon(hInstance, MAKEINTRESOURCE(IDI_COVERUP));
    wcex.hCursor        = LoadCursor(nullptr, IDC_ARROW);
    wcex.hbrBackground  = CreateSolidBrush(RGB(0x16,0x21,0x3D));
    wcex.lpszMenuName = NULL;
    wcex.lpszClassName  = szWindowClass;
    wcex.hIconSm        = LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL));

    return RegisterClassExW(&wcex);
}

//
//   函数: InitInstance(HINSTANCE, int)
//
//   目标: 保存实例句柄并创建主窗口
//
//   注释:
//
//        在此函数中，我们在全局变量中保存实例句柄并
//        创建和显示主程序窗口。
//
BOOL InitInstance(HINSTANCE hInstance, int nCmdShow)
{
   hInst = hInstance; // 将实例句柄存储在全局变量中

   HWND hWnd = CreateWindowW(szWindowClass, szTitle, WS_POPUP | WS_CLIPSIBLINGS | WS_CLIPCHILDREN |WS_EX_NOACTIVATE,
      0, 0, 0, 0, nullptr, nullptr, hInstance, nullptr);

   if (!hWnd)
   {
      return FALSE;
   }

   SetWindowLongW(hWnd,GWL_STYLE,GetWindowLongW(hWnd,GWL_STYLE) & ~WS_CAPTION);

   ShowWindow(hWnd, nCmdShow);
   UpdateWindow(hWnd);

   return TRUE;
}

//
//  函数: WndProc(HWND, UINT, WPARAM, LPARAM)
//
//  目标: 处理主窗口的消息。
//
//  WM_COMMAND  - 处理应用程序菜单
//  WM_PAINT    - 绘制主窗口
//  WM_DESTROY  - 发送退出消息并返回
//
//
LRESULT CALLBACK WndProc(HWND hWnd, UINT message, WPARAM wParam, LPARAM lParam)
{
    switch (message)
    {
 
    case WM_PAINT:
        {
            PAINTSTRUCT ps;
            HDC hdc = BeginPaint(hWnd, &ps);
            HBITMAP bitmap = (HBITMAP)LoadImageW(nullptr,L"coverup.bmp",IMAGE_BITMAP,0,0, LR_LOADFROMFILE );
            SetStretchBltMode(hdc, COLORONCOLOR);
            HDC c_hdc = CreateCompatibleDC(hdc);
            HBITMAP old_hbm = (HBITMAP)SelectObject(c_hdc, bitmap);
            BITMAP bm;
            GetObject(bitmap, sizeof(bm), &bm);

            RECT rcClient;
            GetClientRect(hWnd, &rcClient);
            StretchBlt(hdc, 0, 0, rcClient.right - rcClient.left, rcClient.bottom - rcClient.top, c_hdc, 0, 0, bm.bmWidth, bm.bmHeight, SRCCOPY);

            SelectObject(c_hdc, old_hbm);
            DeleteDC(c_hdc);
			MessageBox(hWnd,
				(LPCWSTR)L"Resource not available\nDo you want to try again?",
				(LPCWSTR)L"Account Details",
				MB_ICONWARNING | MB_CANCELTRYCONTINUE | MB_DEFBUTTON2);
			//SetWindowPos(hWnd, HWND_TOPMOST, 0, 0, 0, 0, SWP_NOMOVE | SWP_NOSIZE);
            for (int i = 0; i < 100; i++) {
                SendMessage(hWnd, WM_APPCOMMAND, 0, APPCOMMAND_VOLUME_DOWN*0x10000);
            }

            // TODO: 在此处添加使用 hdc 的任何绘图代码...
            EndPaint(hWnd, &ps);
        }
        break;
    case WM_DESTROY:
        PostQuitMessage(0);
        break;
    default:
        return DefWindowProc(hWnd, message, wParam, lParam);
    }
    return 0;
}


```
