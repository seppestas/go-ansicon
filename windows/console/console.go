// +build windows

package console

import (
	"fmt"
	"github.com/bitbored/go-ansicon/windows/api"
	"github.com/bitbored/go-ansicon/windows/color"
)

var fg_colors = []uint16{
	0,
	0,
	foreground_red,
	foreground_green,
	foreground_red | foreground_green,
	foreground_blue,
	foreground_red | foreground_blue,
	foreground_green | foreground_blue,
	foreground_red | foreground_green | foreground_blue}

var bg_colors = []uint16{
	0,
	0,
	background_red,
	background_green,
	background_red | background_green,
	background_blue,
	background_red | background_blue,
	background_green | background_blue,
	background_red | background_green | background_blue}

const (
	foreground_blue      = uint16(0x0001)
	foreground_green     = uint16(0x0002)
	foreground_red       = uint16(0x0004)
	foreground_intensity = uint16(0x0008)
	background_blue      = uint16(0x0010)
	background_green     = uint16(0x0020)
	background_red       = uint16(0x0040)
	background_intensity = uint16(0x0080)
	underscore           = uint16(0x8000)
	reverse              = uint16(0x4000)

	foreground_mask = foreground_blue | foreground_green | foreground_red | foreground_intensity
	background_mask = background_blue | background_green | background_red | background_intensity
)

func Reset() {
	winAPI.SetConsoleTextAttribute(winAPI.StdOut, winAPI.InitScreenInfo.WAttributes)
}

func SetAttributes(fg color.Color, fgBright bool, bg color.Color, bgBright, inversed, underlined bool) {
	cbufinfo := winAPI.GetConsoleScreenBufferInfo(winAPI.StdOut)
	if cbufinfo == nil { // No console info - Ex: stdout redirection
		return
	}
	attr := cbufinfo.WAttributes

	if fg != color.None {
		attr = attr & ^foreground_mask | fg_colors[fg]
		if fgBright {
			attr |= foreground_intensity
		}
	}

	if bg != color.None {
		attr = attr & ^background_mask | bg_colors[bg]
		if bgBright {
			attr |= background_intensity
		}
	}

	if inversed {
		attr |= reverse
		// Though Windows' console supports COMMON_LVB_REVERSE_VIDEO,
		// it seems to be buggy. So we must simulate it.
		attr = (attr & underscore) | ((attr & 0x00f0) >> 4) | ((attr & 0x000f) << 4)
	}

	if underlined {
		attr |= underscore
	}

	// fmt.Printf("attr: %b\n", attr)
	if !winAPI.SetConsoleTextAttribute(winAPI.StdOut, attr) {
		fmt.Println("Error happend :(")
	}
}

func SetTitle(title string) {
	winAPI.SetConsoleTitle(title)
}
