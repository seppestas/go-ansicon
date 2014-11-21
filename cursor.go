// +build windows

package cursor

func Up(n int) {
	// if (es_argc == 0) es_argv[es_argc++] = 1; // ESC[A == ESC[1A
	// if (es_argc != 1) return;
	// Pos.Y = CUR.Y - es_argv[0];
	// if (Pos.Y < TOP) Pos.Y = TOP;
	// Pos.X = CUR.X;
	// SetConsoleCursorPosition( hConOut, Pos );
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

func SetPosition(n, m int) {
	// if (es_argc == 0)
	// 	es_argv[es_argc++] = 1; // ESC[H == ESC[1;1H
	// if (es_argc == 1)
	// 	es_argv[es_argc++] = 1; // ESC[#H == ESC[#;1H
	// if (es_argc > 2) return;
	// Pos.X = es_argv[1] - 1;
	// if (Pos.X < LEFT) Pos.X = LEFT;
	// if (Pos.X > RIGHT) Pos.X = RIGHT;
	// Pos.Y = es_argv[0] - 1;
	// if (Pos.Y < TOP) Pos.Y = TOP;
	// if (Pos.Y > BOTTOM) Pos.Y = BOTTOM;
	// SetConsoleCursorPosition( hConOut, Pos );
}

func ForwardTabs(n, m int) {
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

func SetRow(n, int) {
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
