// +build windows

package console

import "github.com/bitbored/go-ansicon/windows/api"

func SetTitle(title string) {
	winAPI.SetConsoleTitle(title)
}
