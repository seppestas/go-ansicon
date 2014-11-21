// +build windows

package winAPI

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetStdHandle               = kernel32.NewProc("GetStdHandle")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")

	HStdout        uintptr
	InitScreenInfo *Console_screen_buffer_info
)

func SetConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		hConsoleOutput,
		uintptr(wAttributes))
	return ret != 0
}

type coord struct {
	X, Y int16
}

type small_rect struct {
	Left, Top, Right, Bottom int16
}

type Console_screen_buffer_info struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            small_rect
	DwMaximumWindowSize coord
}

func GetConsoleScreenBufferInfo(hConsoleOutput uintptr) *Console_screen_buffer_info {
	var csbi Console_screen_buffer_info
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		hConsoleOutput,
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}

const (
	std_output_handle = uint32(-11 & 0xFFFFFFFF)
)

func init() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle = kernel32.NewProc("GetStdHandle")
	HStdout, _, _ = procGetStdHandle.Call(uintptr(std_output_handle))
	InitScreenInfo = GetConsoleScreenBufferInfo(HStdout)
	syscall.LoadDLL("")
}
