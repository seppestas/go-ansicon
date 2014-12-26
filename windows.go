// +build windows

package ansicon

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/bitbored/go-ansicon/controls/csi"
	"github.com/bitbored/go-ansicon/controls/esc"
	"github.com/bitbored/go-ansicon/controls/osc"
	"io"
	"os"
	"strconv"
)

func convert(out io.Writer) (w io.Writer) {
	r, w := io.Pipe()

	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			if scanner.Text()[0] == '\x1b' {
				scanner.Scan()
				handleControl(scanner.Text()[0], scanner)
			} else {
				out.Write(scanner.Bytes())
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error with the scanner in attached container", err)
		}
	}(r)
	return w
}

func handleControl(control byte, scanner *bufio.Scanner) {
	switch control {
	// Functions using CSI
	case '[':
		scanner.Scan()
		switch scanner.Text()[0] {
		case '?':
			scanner.Scan()
			args := getNumericParameters(scanner)
			csi.HandleDECPrivateModeCommand(scanner.Text()[0], args)
		case '!':
			scanner.Scan()
			csi.HandleTerminalCommand(scanner.Text()[0])
		case '>':
			scanner.Scan()
			args := getNumericParameters(scanner)
			csi.HandleSpecialCommand(scanner.Text()[0], args)
		default:
			args := getNumericParameters(scanner)
			csi.HandleCommand(scanner.Text()[0], args)
		}

	// Operating System Controls
	case ']':
		scanner.Scan()
		ps := getNumericParameters(scanner)
		pt := getTextParameter(scanner)
		osc.HandleCommand(ps[0], pt)

	// Controls beginning with ESC
	case ' ', '#', '%', '(', ')', '*', '+', '-', '.', '/':
		scanner.Scan()
		esc.HandleCharacterSetCommand(control, scanner.Text()[0])
	default:
		esc.HandleCommand(control)
	}
}

// Reads numeric parameters from the scanner
func getNumericParameters(scanner *bufio.Scanner) []int {
	args := make([]int, 0)
	arg := 0
	gotArg := false

	for {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			if gotArg {
				arg = arg*10 + n
			} else {
				arg = n
				gotArg = true
			}
		} else { // Not an int
			if gotArg {
				args = append(args, arg)
				arg = 0
				gotArg = false
			}
			if scanner.Text()[0] != ';' {
				return args
			}
		}
		scanner.Scan()
	}
}

// Reads the ; seperated arguments from the scanner
func getTextParameter(scanner *bufio.Scanner) string {
	var buffer bytes.Buffer
	for {
		if scanner.Text()[0] == '\x07' || scanner.Text()[0] == '\x9c' {
			return buffer.String()
		}
		// fmt.Printf(" -0x%x- ", scanner.Text()[0])
		// fmt.Print(scanner.Text())
		buffer.WriteString(scanner.Text())
		scanner.Scan()
	}

}
