package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) int {
	n := 0
	sign := 1

	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c-'0')
	}
	return n * sign
}

func main() {
	args := os.Args[1:]
	upper := false

	// Check for --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	// Process each argument
	for _, arg := range args {
		n := atoi(arg)
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		// Calculate the corresponding letter
		var base rune = 'a' - 1
		if upper {
			base = 'A' - 1
		}
		z01.PrintRune(base + rune(n))
	}

	// Print newline at the end if there were arguments
	if len(args) > 0 {
		z01.PrintRune('\n')
	}
}
