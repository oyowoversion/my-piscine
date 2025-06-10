package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	// Skip the first argument (program name) and print the rest
	for _, arg := range os.Args[1:] {
		printStr(arg)
	}
}
