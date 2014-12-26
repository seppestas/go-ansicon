// +build windows

package display

import "github.com/bitbored/go-ansicon/windows/color"

func createColor(code int) color.Color {
	return color.Color(code) + color.Black
}

func SelectGraphicRendition(args []int) {
	fg_bright := false
	bg_bright := false
	fg_color := color.None
	bg_color := color.None
	for _, arg := range args {
		switch {

		case arg == 0:
			color.ResetColor()
			fg_bright = false
			bg_bright = false
			fg_color = color.None
			bg_color = color.None

		case arg == 1:
			fg_bright = true

		case arg == 4:
			bg_bright = true

		case 30 <= arg && arg <= 37:
			fg_color = createColor(arg - 30)

		case arg == 39:
			fg_color = color.None

		case 40 <= arg && arg <= 47:
			bg_color = createColor(arg - 40)

		case arg == 49:
			fg_color = color.None
		}
	}
	color.ChangeColor(fg_color, fg_bright, bg_color, bg_bright)
}
