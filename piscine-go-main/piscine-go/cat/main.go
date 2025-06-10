package main

import (
	"io"
	"os"
)

func printError(msg string) {
	os.Stderr.Write([]byte(msg))
}

func printFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		printError("ERROR: " + err.Error() + "\n")
		return err
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	return err
}

func printStdin() error {
	_, err := io.Copy(os.Stdout, os.Stdin)
	return err
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		err := printStdin()
		if err != nil {
			os.Exit(1)
		}
		return
	}

	for _, filename := range args {
		if err := printFile(filename); err != nil {
			os.Exit(1) // exit immediately on first error
		}
	}
}
