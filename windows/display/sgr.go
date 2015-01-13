// +build windows

package display

import (
	"github.com/bitbored/go-ansicon/windows/color"
	"github.com/bitbored/go-ansicon/windows/console"
)

func createColor(code int) color.Color {
	return color.Color(code) + color.Black
}

func SelectGraphicRendition(args []int) {
	fg_bright := false
	bg_bright := false
	fg_color := color.None
	bg_color := color.None
	inversed := false
	underlined := false

	if len(args) == 0 {
		console.Reset()
		return
	}

	for _, arg := range args {
		switch arg {
		case 0:
			console.Reset()
			fg_bright = false
			bg_bright = false
			fg_color = color.None
			bg_color = color.None

		case 1:
			fg_bright = true

		case 4:
			underlined = true

		case 5:
			bg_bright = true

		case 7:
			inversed = true

		// Following the xterm spec for 20-30
		case 22:
			fg_bright = false
			bg_bright = false

		case 24:
			underlined = false

		case 25:
			bg_bright = false

		case 27:
			inversed = false

		case 39:
			fg_color = color.None

		case 49:
			fg_color = color.None

		default:
			if 30 <= arg && arg <= 37 {
				fg_color = createColor(arg - 30)
			} else if 40 <= arg && arg <= 47 {
				bg_color = createColor(arg - 40)
			} else if 90 <= arg && arg <= 97 {
				fg_color = createColor(arg - 90)
				fg_bright = true
			} else if 100 <= arg && arg <= 107 {
				bg_color = createColor(arg - 100)
				bg_bright = true
			}
		}
	}
	console.SetAttributes(fg_color, fg_bright, bg_color, bg_bright, inversed, underlined)
}
