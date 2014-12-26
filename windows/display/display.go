// +build windows

package display

import "github.com/bitbored/go-ansicon/windows/api"

var (
	screenBufferInfo *winAPI.ConsoleScreenBufferInfo
	screenTop        int
)

func init() {
	screenTop = -1
}

func Erase(n int) {
	screenBufferInfo = winAPI.GetConsoleScreenBufferInfo(winAPI.StdOut)

	width := screenBufferInfo.DwSize.X
	cur := screenBufferInfo.DwCursorPosition
	top := screenBufferInfo.SrWindow.Top
	bottom := screenBufferInfo.SrWindow.Bottom
	screenTop := -1
	left := 0
	right := width - 1
	var pos winAPI.Coord

	switch n {
	case 0: // Erase Below
		length := int((bottom-cur.Y)*width + width - cur.X)
		fillBlank(length, cur)
	case 1: // Erase Above
		pos.X = 0
		pos.Y = top
		length := int((cur.Y-top)*width + cur.X + 1)
		fillBlank(length, pos)
	case 2: // Erase All
		if int(top) != screenTop || bottom == screenBufferInfo.DwSize.Y-1 {
			// Rather than clearing the existing window, make the current
			// line the new top of the window (assuming this is the first
			// thing a program does).
			size := bottom - top
			if cur.Y+size < screenBufferInfo.DwSize.Y {
				top = cur.Y
				bottom = top + size
			} else {
				bottom = screenBufferInfo.DwSize.Y - 1
				top = bottom - size
				var rect winAPI.SmallRect
				rect.Left = int16(left)
				rect.Right = right
				rect.Top = cur.Y - 1
				pos.X, pos.Y = 0, 0
				var charInfo winAPI.CharInfo
				charInfo.Char[1] = ' '
				charInfo.Attributes = screenBufferInfo.WAttributes
				winAPI.ScrollConsoleScreenBuffer(winAPI.StdOut, rect, winAPI.SmallRect{}, pos, charInfo)
			}
			winAPI.SetConsoleWindowInfo(winAPI.StdOut, true, &screenBufferInfo.SrWindow)
			screenTop = int(top)
		}
		pos.X = int16(left)
		pos.Y = top
		length := (bottom - top + 1) * width
		fillBlank(int(length), pos)
	}
}

func fillBlank(length int, pos winAPI.Coord) {
	winAPI.FillConsoleOutputCharacter(winAPI.StdOut, ' ', uint32(length), pos)
	winAPI.FillConsoleOutputAttribute(winAPI.StdOut, screenBufferInfo.WAttributes, uint32(length), pos)
}

func EraseInLine(n int) {

}

func ScrollUp(n int) {

}

func ScrollDown(n int) {

}

func RepeatCharacter(n int) {

}
