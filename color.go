// +build windows

package display

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

func resetColor() {
	if initScreenInfo == nil { // No console info - Ex: stdout redirection
		return
	}
	setConsoleTextAttribute(hStdout, initScreenInfo.WAttributes)
}

func changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	attr := uint16(0)
	if fg == None || bg == None {
		cbufinfo := getConsoleScreenBufferInfo(hStdout)
		if cbufinfo == nil { // No console info - Ex: stdout redirection
			return
		}
		attr = getConsoleScreenBufferInfo(hStdout).WAttributes
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

	setConsoleTextAttribute(hStdout, attr)
}
