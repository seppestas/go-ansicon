// +build !windows

package ansicon

import "io"

func convert(w io.Writer) io.Writer {
	return w
}
