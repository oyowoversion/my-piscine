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
	// Get arguments (excluding program name)
	args := os.Args[1:]

	// Print in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		printStr(args[i])
	}
}
