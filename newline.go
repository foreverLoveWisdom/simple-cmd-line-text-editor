package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func editFile(filename string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type your text below (type 'SAVE' to save and exit):")

	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			break
		}

		// Check for the 'SAVE' command to exit the loop
		if strings.TrimSpace(line) == "SAVE" {
			break
		}

		lines = append(lines, line)
	}

	if len(lines) == 0 {
		fmt.Println("No content to save.")
		return
	}

	err := os.WriteFile(filename, []byte(strings.Join(lines, "")), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File saved successfully!")
}

func main() {
	filename := "example.txt"
	editFile(filename)
}
