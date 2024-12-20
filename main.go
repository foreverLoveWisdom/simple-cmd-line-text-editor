package main

import (
	"bufio"
	"fmt"
	"os"
)

func editFile(filename string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type your text below (type 'SAVE' to save and exit):")

	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if line == "SAVE\n" {
			break
		}

		lines = append(lines, line)
	}

	err := os.WriteFile(filename, []byte(fmt.Sprint(lines)), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File saved successfully!")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Printf("Editing file: %s\n", filename)

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Printf("\nContents of %s:\n%s\n", filename, content)
	editFile(filename)
}
