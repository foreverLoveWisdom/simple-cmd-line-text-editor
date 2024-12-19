package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now() // Start timing

	// Simulate reading from an I/O-bound source (the file)
	fmt.Println("Reading from file...")
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the content of the file
	fmt.Println(string(data))

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("File read completed in %v seconds.\n", elapsed.Seconds())
}
