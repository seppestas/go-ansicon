// +build windows

package cursor

import "github.com/bitbored/go-ansicon/windows/api"

func Up(n int) {
	screenBufferInfo := winAPI.GetConsoleScreenBufferInfo(winAPI.StdOut)

	cur := screenBufferInfo.DwCursorPosition
	top := screenBufferInfo.SrWindow.Top

	var pos winAPI.Coord
	pos.Y = cur.Y - int16(n)
	if pos.Y < top {
		pos.Y = top
	}
	pos.X = cur.X
	winAPI.SetConsoleCursorPosition(winAPI.StdOut, pos)
}

func Down(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[B == ESC[1B
	// if (es_argc != 1) return;
	// Pos.Y = CUR.Y + es_argv[0];
	// if (Pos.Y > BOTTOM) Pos.Y = BOTTOM;
	// Pos.X = CUR.X;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func Forward(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[C == ESC[1C
	// if (es_argc != 1) return;
	// Pos.X = CUR.X + es_argv[0];
	// if (Pos.X > RIGHT) Pos.X = RIGHT;
	// Pos.Y = CUR.Y;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func Back(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[D == ESC[1D
	// if (es_argc != 1) return;
	// Pos.X = CUR.X - es_argv[0];
	// if (Pos.X < LEFT) Pos.X = LEFT;
	// Pos.Y = CUR.Y;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func NextLine(n int) {
	// 	if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[E == ESC[1E
	// 	if (es_argc != 1) return;
	// 	Pos.Y = CUR.Y + es_argv[0];
	// 	if (Pos.Y > BOTTOM) Pos.Y = BOTTOM;
	// 	Pos.X = LEFT;
	// 	SetConsoleCursorPosition( hConOut, Pos );
}

func PreviousLine(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[F == ESC[1F
	// if (es_argc != 1) return;
	// Pos.Y = CUR.Y - es_argv[0];
	// if (Pos.Y < TOP) Pos.Y = TOP;
	// Pos.X = LEFT;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func HorizontalAbsolute(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[G == ESC[1G
	// if (es_argc != 1) return;
	// Pos.X = es_argv[0] - 1;
	// if (Pos.X > RIGHT) Pos.X = RIGHT;
	// if (Pos.X < LEFT) Pos.X = LEFT;
	// Pos.Y = CUR.Y;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func SetPosition(args []int) {
	screenBufferInfo := winAPI.GetConsoleScreenBufferInfo(winAPI.StdOut)

	width := screenBufferInfo.DwSize.X
	top := screenBufferInfo.SrWindow.Top
	bottom := screenBufferInfo.SrWindow.Bottom

	left := int16(0)
	right := int16(width - 1)
	if len(args) > 2 {
		return
	}

	pos := winAPI.Coord{1, 1}

	if len(args) > 0 {
		pos.X = int16(args[0]) // ESC[#H == ESC[#;1H
	}
	if len(args) > 1 {
		pos.Y = int16(args[1])
	}

	if pos.X < left {
		pos.X = left
	} else if pos.X > right {
		pos.X = right
	}
	if pos.Y < top {
		pos.Y = top
	} else if pos.Y > bottom {
		pos.Y = bottom
	}
	winAPI.SetConsoleCursorPosition(winAPI.StdOut, pos)
}

func ForwardTabs(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[I == ESC[1I
	// if (es_argc != 1) return;
	// Pos.Y = CUR.Y;
	// Pos.X = (CUR.X & -8) + es_argv[0] * 8;
	// if (Pos.X > RIGHT) Pos.X = RIGHT;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func BackTabs(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[Z == ESC[1Z
	// if (es_argc != 1) return;
	// Pos.Y = CUR.Y;
	// if ((CUR.X & 7) == 0)
	// 	 Pos.X = CUR.X - es_argv[0] * 8;
	// else
	// 	 Pos.X = (CUR.X & -8) - (es_argv[0] - 1) * 8;
	// if (Pos.X < LEFT) Pos.X = LEFT;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func SetRow(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[d == ESC[1d
	// if (es_argc != 1) return;
	// Pos.Y = es_argv[0] - 1;
	// if (Pos.Y < TOP) Pos.Y = TOP;
	// if (Pos.Y > BOTTOM) Pos.Y = BOTTOM;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func Move(n, m int) {

}

func SavePosition() {
	// if (es_argc != 0) return;
	// pState->SavePos.X = CUR.X;
	// pState->SavePos.Y = CUR.Y - TOP;

}

func RestorePosition() {
	// if (es_argc != 0) return;
	// Pos.X = pState->SavePos.X;
	// Pos.Y = pState->SavePos.Y + TOP;
	// if (Pos.X > RIGHT) Pos.X = RIGHT;
	// if (Pos.Y > BOTTOM) Pos.Y = BOTTOM;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func Hide() {
	// GetConsoleCursorInfo( hConOut, &CursInfo );
	// CursInfo.bVisible = (suffix == 'h');
	// SetConsoleCursorInfo( hConOut, &CursInfo );
}

func Show() {
	// GetConsoleCursorInfo( hConOut, &CursInfo );
	// CursInfo.bVisible = (suffix == 'h');
	// SetConsoleCursorInfo( hConOut, &CursInfo );
}
