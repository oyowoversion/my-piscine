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
	// The first element of os.Args is the program name
	programName := os.Args[0]
	// Extract just the executable name (without path)
	// We need to find the last '/' or '\' in the path
	start := 0
	for i, c := range programName {
		if c == '/' || c == '\\' {
			start = i + 1
		}
	}
	programName = programName[start:]
	printStr(programName)
}
