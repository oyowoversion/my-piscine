package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Copy file contents to stdout
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
}
