// +build windows

package csi

import (
	"github.com/bitbored/go-ansicon/windows/cursor"
	"github.com/bitbored/go-ansicon/windows/display"
	"github.com/bitbored/go-ansicon/windows/terminal"
)

/*
Control Sequence Inducer
This package contains logic for functions using the CSI control
*/
func HandleCommand(command byte, args []int) {
	n := 0
	if len(args) > 0 {
		n = args[0]
	}

	// fmt.Printf(" { Command : %c, args: %v }", command, args)
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
		cursor.SetPosition(args)
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
		display.ScrollDown(n)
	case 'b':
		display.RepeatCharacter(n)
	case 'd':
		cursor.SetRow(n)
	case 'n':
		// statusReport(n)
	case 't':
		terminal.ReportTitle(n)
	case 'h':
		terminal.SetMode(n)
	case 'l':
		terminal.ResetMode(n)
	case 'm':
		display.SelectGraphicRendition(args)
	case 'i':
		// Set AUX port, ignored
	}
}

func HandleDECPrivateModeCommand(command byte, args []int) {
	switch command {
	case 'J':
		// DECSED
	case 'K':
		// DECSEL
	case 'h':
		// DECSET
		// ...
	}
}

func HandleTerminalCommand(command byte) {
	if command == 'p' {
		// DECSTR
	}
}

func HandleSpecialCommand(command byte, args []int) {
	// Not (yet) implemented
}
