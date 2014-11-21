// +build !windows

package ansicon

import "io"

func convert(input io.Writer) io.Writer {
	return input
}
