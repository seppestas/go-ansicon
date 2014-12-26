package osc

import (
	"github.com/bitbored/go-ansicon/windows/console"
)

func HandleCommand(command int, p string) {
	switch command {
	case 0:
	case 1:
	case 2:
		console.SetTitle(p)
	}
}
