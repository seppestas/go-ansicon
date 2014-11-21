// +build windows

package xterm

func ReportTitle(n int) {
	// if (es_argc != 1) return;
	// if (es_argv[0] == 21)	// ESC[21t Report xterm window's title
	// {
	// 	TCHAR buf[MAX_PATH*2];
	// 	DWORD len = GetConsoleTitle( buf+3, lenof(buf)-3-2 );
	// 	// Too bad if it's too big or fails.
	// 	buf[0] = ESC;
	// 	buf[1] = ']';
	// 	buf[2] = 'l';
	// 	buf[3+len] = ESC;
	// 	buf[3+len+1] = '\\';
	// 	buf[3+len+2] = '\0';
	// 	SendSequence( buf );
	// }
}

func SetMode(n int) {
	// if (es_argc == 1 && es_argv[0] == 3)
	// 	pState->crm = TRUE;

}
func ResetMode(n int) {

}
