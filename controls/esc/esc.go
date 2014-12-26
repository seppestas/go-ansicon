package esc

func HandleCharacterSetCommand(command, option byte) {
	switch command {
	case ' ':
		switch option {
		case 'F':
		case 'G':
		case 'L':
		case 'M':
		case 'N':
		}
	case '#':
		switch option {
		case '3':
		case '4':
		case '5':
		case '6':
		case '8':
		}
	case '%':
		switch option {
		case '@':
		case 'G':
		}
	case '(':
		designateCharacterSet(0, option)
	case ')', '-':
		designateCharacterSet(1, option)
	case '*', '.':
		designateCharacterSet(2, option)
	case '+', '/':
		designateCharacterSet(3, option)
	}

}

func designateCharacterSet(index int, c byte) {

}

func HandleCommand(command byte) {

}
