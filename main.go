package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s command [args]\n", os.Args[0])
	fmt.Fprint(os.Stderr, "commands:\n")
	for _, c := range Commands {
		fmt.Fprintf(os.Stderr, "  %-10s - %s\n", c.Name(), c.Help())
	}
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	for _, c := range Commands {
		if os.Args[1] == c.Name() {
			fmt.Println(c.Run(os.Args[1:]))
		}
	}
}