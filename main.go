package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileEditor struct {
	Filename string
}

func (fe *FileEditor) Load() (*string, error) {
	content, err := os.ReadFile(fe.Filename)
	if err != nil {
		return nil, err
	}
	contentStr := string(content)
	return &contentStr, nil
}

func (fe *FileEditor) Save(content *string) error {
	return os.WriteFile(fe.Filename, []byte(*content), 0644)
}

func (fe *FileEditor) Edit() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type your text below (type 'SAVE' to save and exit):")

	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if line == "SAVE\n" {
			break
		}

		lines = append(lines, line)
	}

	content := ""
	for _, line := range lines {
		content += line
	}

	return fe.Save(&content)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	fe := &FileEditor{Filename: os.Args[1]}
	fmt.Printf("Editing file: %s\n", fe.Filename)

	content, err := fe.Load()
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Printf("\nContents of %s:\n%s\n", fe.Filename, *content)
	}

	if err := fe.Edit(); err != nil {
		fmt.Println("Error during editing:", err)
	}
}
