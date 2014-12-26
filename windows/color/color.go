// +build windows

package color

import "github.com/bitbored/go-ansicon/windows/api"

type Color int

const (
	// No change of color
	None = Color(iota)
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
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

	foreground_mask = foreground_blue | foreground_green | foreground_red | foreground_intensity
	background_mask = background_blue | background_green | background_red | background_intensity
)

func ResetColor() {
	winAPI.SetConsoleTextAttribute(winAPI.StdOut, winAPI.InitScreenInfo.WAttributes)
}

func ChangeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	attr := uint16(0)
	if fg == None || bg == None {
		cbufinfo := winAPI.GetConsoleScreenBufferInfo(winAPI.StdOut)
		if cbufinfo == nil { // No console info - Ex: stdout redirection
			return
		}
		attr = winAPI.GetConsoleScreenBufferInfo(winAPI.StdOut).WAttributes
	} // if

	if fg != None {
		attr = attr & ^foreground_mask | fg_colors[fg]
		if fgBright {
			attr |= foreground_intensity
		} // if
	} // if

	if bg != None {
		attr = attr & ^background_mask | bg_colors[bg]
		if bgBright {
			attr |= background_intensity
		} // if
	} // if

	winAPI.SetConsoleTextAttribute(winAPI.StdOut, attr)
}
