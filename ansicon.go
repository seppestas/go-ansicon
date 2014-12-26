package ansicon

import "io"

// On Windows: Converts the ANSI characters written to input to the corresponding Winwows API calls and writes data
// without escape characters to outpt using a Go routine
// On Non-Windows: just returns the given io.Writer
func Convert(output io.Writer) io.Writer {
	return convert(output)
}
