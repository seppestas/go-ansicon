// +build windows

package winAPI

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"unsafe"
)

/*
#include <stdlib.h>

union Char {
	unsigned short UnicodeChar;
	char           AsciiChar;
} Char;
*/
import "C"

const (
	STDOUT_HANDLE = uint32(-11 & 0xFFFFFFFF)
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetStdHandle               = kernel32.NewProc("GetStdHandle")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	procSetConsoleTitle            = kernel32.NewProc("SetConsoleTitleW")
	procFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterA")
	procFillConsoleOutputAttribute = kernel32.NewProc("FillConsoleOutputAttribute")
	procScrollConsoleScreenBuffer  = kernel32.NewProc("ScrollConsoleScreenBufferA")
	procSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
	procSetConsoleWindowInfo       = kernel32.NewProc("SetConsoleWindowInfo")
	procGetLastError               = kernel32.NewProc("GetLastError")

	StdOut         uintptr
	InitScreenInfo *ConsoleScreenBufferInfo
)

type Coord struct {
	X, Y int16
}

func (c Coord) toWord() uint32 {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, c)
	if err != nil {
		panic(err)
	}
	var w uint32
	err = binary.Read(buf, binary.LittleEndian, &w)
	if err != nil {
		panic(err)
	}
	return w
}

type SmallRect struct {
	Left, Top, Right, Bottom int16
}

type ConsoleScreenBufferInfo struct {
	DwSize              Coord
	DwCursorPosition    Coord
	WAttributes         uint16
	SrWindow            SmallRect
	DwMaximumWindowSize Coord
}

type CharInfo struct {
	Char       C.union_Char
	Attributes uint16
}

func GetConsoleScreenBufferInfo(hConsoleOutput uintptr) *ConsoleScreenBufferInfo {
	var csbi ConsoleScreenBufferInfo
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		hConsoleOutput,
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}

func SetConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		hConsoleOutput,
		uintptr(wAttributes))
	return ret != 0
}

func SetConsoleTitle(title string) bool {
	strptr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return false
	}
	ret, _, _ := procSetConsoleTitle.Call(uintptr(unsafe.Pointer(strptr)))
	return (ret != 0)
}

func FillConsoleOutputCharacter(hConsoleOutput uintptr, cCharacter byte, nLength uint32, dwWriteCoord Coord) (numberOfCharsWritten int, err bool) {
	var charsWritten uint32
	ret, _, _ := procFillConsoleOutputCharacter.Call(
		hConsoleOutput,
		uintptr(cCharacter),
		uintptr(nLength),
		uintptr(dwWriteCoord.toWord()),
		uintptr(unsafe.Pointer(&charsWritten)))
	return int(charsWritten), (ret != 0)
}

func FillConsoleOutputAttribute(hConsoleOutput uintptr, wAttributes uint16, nLength uint32, dwWriteCoord Coord) (numberOfCharsWritten int, errors bool) {
	var charsWritten int
	ret, _, _ := procFillConsoleOutputAttribute.Call(
		hConsoleOutput,
		uintptr(wAttributes),
		uintptr(nLength),
		uintptr(dwWriteCoord.toWord()),
		uintptr(unsafe.Pointer(&charsWritten)))
	return charsWritten, (ret != 0)
}

func ScrollConsoleScreenBuffer(hConsoleOutput uintptr, lpScrollRectangle, lpClipRectangle SmallRect, dwDestinationOrigin Coord, lpFill CharInfo) bool {
	ret, _, _ := procScrollConsoleScreenBuffer.Call(
		hConsoleOutput,
		uintptr(unsafe.Pointer(&lpScrollRectangle)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&dwDestinationOrigin)),
		uintptr(unsafe.Pointer(&lpFill)))
	return ret != 0
}

func SetConsoleCursorPosition(hConsoleOutput uintptr, dwCursorPosition Coord) bool {
	ret, _, _ := procSetConsoleCursorPosition.Call(
		hConsoleOutput,
		uintptr(dwCursorPosition.toWord()))
	return ret != 0
}

func SetConsoleWindowInfo(hConsoleOutput uintptr, bAbsolute bool, lpConsoleWindow *SmallRect) bool {
	abs := 0
	if bAbsolute {
		abs = 1
	}
	ret, _, _ := procSetConsoleWindowInfo.Call(
		hConsoleOutput,
		uintptr(abs),
		uintptr(unsafe.Pointer(lpConsoleWindow)))
	return ret != 0
}

func GetLastError() uint16 {
	ret, _, _ := procGetLastError.Call()
	return uint16(ret)
}

func init() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle = kernel32.NewProc("GetStdHandle")
	StdOut, _, _ = procGetStdHandle.Call(uintptr(STDOUT_HANDLE))
	InitScreenInfo = GetConsoleScreenBufferInfo(StdOut)
	syscall.LoadDLL("")
}
