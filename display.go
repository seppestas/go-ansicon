// +build windows

package display

func createColor(code int) Color {
	return Color(code) + Black
}

func SelectGraphicRendition(args []int) {
	fg_bright := false
	fg_color := None
	bg_bright := false
	bg_color := None
	for _, arg := range args {
		switch {

		case arg == 0:
			resetColor()
			fg_bright = false
			fg_color = None
			bg_bright = false
			bg_color = None

		case arg == 1:
			fg_bright = true

		case arg == 4:
			bg_bright = true

		case 30 <= arg <= 37:
			fg_color = createColor(arg - 30)

		case arg == 39:
			fg_color = None

		case 40 <= arg <= 47:
			bg_color = createColor(arg - 40)

		case arg == 49:
			fg_color = None
		}
	}
	changeColor(fg_color, fg_bright, bg_color, bg_bright)
}

func Erase(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 0; // ESC[J == ESC[0J
	// if (es_argc != 1) return;
	// switch (es_argv[0])
	// {
	// 	case 0:		// ESC[0J erase from cursor to end of display
	// 		len = (BOTTOM - CUR.Y) * WIDTH + WIDTH - CUR.X;
	// 		FillBlank( len, CUR );
	// 	return;

	// 	case 1:		// ESC[1J erase from start to cursor.
	// 		Pos.X = 0;
	// 		Pos.Y = TOP;
	// 		len   = (CUR.Y - TOP) * WIDTH + CUR.X + 1;
	// 		FillBlank( len, Pos );
	// 	return;

	// 	case 2:		// ESC[2J Clear screen and home cursor
	// 		if (TOP != screen_top || BOTTOM == Info.dwSize.Y - 1)
	// 		{
	// 			// Rather than clearing the existing window, make the current
	// 			// line the new top of the window (assuming this is the first
	// 			// thing a program does).
	// 			int range = BOTTOM - TOP;
	// 			if (CUR.Y + range < Info.dwSize.Y)
	// 			{
	// 	TOP = CUR.Y;
	// 	BOTTOM = TOP + range;
	// 			}
	// 			else
	// 			{
	// 	BOTTOM = Info.dwSize.Y - 1;
	// 	TOP = BOTTOM - range;
	// 	Rect.Left = LEFT;
	// 	Rect.Right = RIGHT;
	// 	Rect.Top = CUR.Y - TOP;
	// 	Rect.Bottom = CUR.Y - 1;
	// 	Pos.X = Pos.Y = 0;
	// 	CharInfo.Char.UnicodeChar = ' ';
	// 	CharInfo.Attributes = Info.wAttributes;
	// 	ScrollConsoleScreenBuffer(hConOut, &Rect, NULL, Pos, &CharInfo);
	// 			}
	// 			SetConsoleWindowInfo( hConOut, TRUE, &WIN );
	// 			screen_top = TOP;
	// 		}
	// 		Pos.X = LEFT;
	// 		Pos.Y = TOP;
	// 		len   = (BOTTOM - TOP + 1) * WIDTH;
	// 		FillBlank( len, Pos );
	// 		// Not technically correct, but perhaps expected.
	// 		SetConsoleCursorPosition( hConOut, Pos );
	// 	return;

	// 	default:
	// 	return;
	// }

}
