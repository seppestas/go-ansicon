// +build windows

package kernel

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetStdHandle               = kernel32.NewProc("GetStdHandle")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")

	hStdout        uintptr
	initScreenInfo *console_screen_buffer_info
)

func setConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
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

type console_screen_buffer_info struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            small_rect
	DwMaximumWindowSize coord
}

func getConsoleScreenBufferInfo(hConsoleOutput uintptr) *console_screen_buffer_info {
	var csbi console_screen_buffer_info
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

	hStdout, _, _ = procGetStdHandle.Call(uintptr(std_output_handle))

	initScreenInfo = getConsoleScreenBufferInfo(hStdout)

	syscall.LoadDLL("")
}
