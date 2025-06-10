package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the target words
	targets := map[string]bool{
		"01":        true,
		"galaxy":    true,
		"galaxy 01": true,
	}

	// Loop through command-line arguments
	for _, arg := range os.Args[1:] {
		if targets[arg] {
			fmt.Println("Alert!!!")
			return
		}
	}
}
