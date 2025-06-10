package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || containsHelp(args) {
		printHelp()
		return
	}

	var inputStr string
	var insertStr string
	var shouldOrder bool

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--order" || arg == "-o":
			shouldOrder = true
		case startsWith(arg, "--insert="):
			insertStr = arg[9:]
		case startsWith(arg, "-i="):
			insertStr = arg[3:]
		case arg != "-o" && arg != "--order" && arg != "-i" && arg != "--insert":
			inputStr = arg
		}
	}

	result := inputStr

	if insertStr != "" {
		result += insertStr
	}

	if shouldOrder {
		result = sortString(result)
	}

	printString(result)
}

func containsHelp(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	helpText := []string{
		"--insert",
		"  -i",
		"\t This flag inserts the string into the string passed as argument.",
		"--order",
		"  -o",
		"\t This flag will behave like a boolean, if it is called it will order the argument.",
	}

	for _, line := range helpText {
		printString(line)
	}
}
