package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text (type 'exit' to quit):")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if line == "exit\n" {
			break
		}

		fmt.Printf("You entered: %s", line)
	}
}
