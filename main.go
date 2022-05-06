package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing argument: contents")
		os.Exit(1)
	}

	qr, err := Create(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error encoding data:", err)
		os.Exit(1)
	}

	Draw(qr, "\x1b[40m", "\x1b[47m")
	os.Stdout.Write([]byte{'\x1b', '[', '0', 'm'})
}
