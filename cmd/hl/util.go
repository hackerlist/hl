package main

import (
	"strings"
	"syscall"
	"unsafe"
)

func wrap(str string, width int) []string {
	var out []string

	lines := strings.Split(str, "\n")

	for _, l := range lines {

		words := strings.Split(l, " ")
		if len(words) == 0 {
			return out
		}

		// current line we are making
		current := words[0]

		// # spaces left before the end
		remaining := width - len(current)

		for _, word := range words[1:] {
			if len(word)+1 > remaining {
				out = append(out, current)
				current = word
				remaining = width - len(word)
			} else {
				current += " " + word
				remaining -= 1 + len(word)
			}
		}

		out = append(out, current)
	}

	return out
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func TerminalWidth() int {
	sizeobj, _ := GetWinsize()
	return int(sizeobj.Col)
}

func GetWinsize() (*winsize, error) {
	ws := new(winsize)

	r1, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(r1) == -1 {
		return nil, err
	}
	return ws, nil
}
