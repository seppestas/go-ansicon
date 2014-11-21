// +build windows

package ansi

import (
	"bufio"
	"bytes"
	"cursor"
	"graphic"
	"io"
	"strconv"
)

const ( // States
	BEGIN
	ESCAPED
	CSI_DONE
	DECTCEM
)

func convert(input io.Writer) (w io.Writer) {
	r, w := io.Pipe()

	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanBytes)
		state := BEGIN
		args := make([]int, 2)
		i := 0
		var buffer bytes.Buffer
		for scanner.Scan() {
			switch state {

			case BEGIN:
				i = 0
				if scanner.Text() == "\x1b" {
					state = ESCAPED
				} else {
					fmt.Printf(scanner.Text())
				}

			case ESCAPED:
				if scanner.Text() == "[" {
					state = CSI_DONE
				} else {
					state = BEGIN
				}

			case CSI_DONE:
				if nPart, err := strconv.Atoi(scanner.Text()); err == nil {
					n = n*10 + nPart
					continue
				} else {
					switch scanner.Text()[0] {
					case ';':
						// Read next arg
						i++
					case 25:
						state = DECTCEM
					default:
						handleCommand(scanner.Text(), args)
						state = BEGIN
					}
				}

			case DECTCEM:
				switch scanner.Text()[0] {
				case 'l':
					cursor.Hide()
				case 'h':
					cursor.Show()
				}
				state = BEGIN
			}

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error with the scanner in attached container", err)
		}
	}(r)
}

func handleCommand(command byte, args []int) {
	n := args[0]
	if len(args) > 1 {
		m := args[1]
	} else {
		m := 0
	}

	switch command {
	case 'A', 'k':
		cursor.Up(n)
	case 'B', 'e':
		cursor.Down(n)
	case 'C', 'a':
		cursor.Forward(n)
	case 'D', 'j':
		cursor.Back(n)
	case 'E':
		cursor.NextLine(n)
	case 'F':
		cursor.PreviousLine(n)
	case 'G', '`':
		cursor.HorizontalAbsolute(n)
	case 'H', 'f':
		cursor.SetPosition(n, m)
	case 'I':
		cursor.ForwardTabs(n)
	case 'Z':
		cursor.BackTabs(n)
	case 'J':
		display.Erase(n)
	case 'K':
		display.EraseInLine(n)
	case 'S':
		display.ScrollUp(n)
	case 'T':
		display.scrollDown(n)
	case 'b':
		display.repeatCharacter(n)
	case 'd':
		cursor.SetRow(n)
	case 's':
		cursor.SavePosition()
	case 'u':
		cursor.RestorePosition()
	case 'n':
		statusReport(n)
	case 't':
		xterm.ReportTitle(n)
	case 'h':
		xterm.SetMode(n)
	case 'l':
		xterm.ResetMode()
	case 'm':
		display.SelectGraphicRendition(args)
	case 'i':
		// Set AUX port, ignored
	}
}

func statusReport(n int) {
	// if (es_argc != 1) return; // ESC[n == ESC[0n -> ignored
	// switch (es_argv[0])
	// {
	// 	case 5:		// ESC[5n Report status
	// 		SendSequence( L"\33[0n" ); // "OK"
	// 	return;

	// 	case 6:		// ESC[6n Report cursor position
	// 	{
	// 		TCHAR buf[32];
	// 		wsprintf( buf, L"\33[%d;%dR", CUR.Y - TOP + 1, CUR.X + 1 );
	// 		SendSequence( buf );
	// 	}
	// 	return;

	// 	default:
	// 	return;
	// }

	// 		case 't':                 // ESC[#t Window manipulation
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
