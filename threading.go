package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

// Struct to hold the response from the API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Function to simulate a task
func task(id int) {
	fmt.Printf("Task %d is starting...\n", id)
	// sleepDuration := rand.Intn(4) + 1 // Sleep between 1 and 4 seconds
	sleepDuration := 0.0000001
	fmt.Printf(
		"Sleep duration for the current task: %d seconds for Task ID: %d\n",
		sleepDuration,
		id,
	)
	time.Sleep(time.Duration(sleepDuration) * time.Second)
	fmt.Printf("Task %d is completed after %d seconds!\n", id, sleepDuration)
}

// Function to perform a real API call
func apiCall(id int) {
	fmt.Printf("API Call for Task %d is starting...\n", id)

	// Make an HTTP GET request to a public API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}
	defer resp.Body.Close()

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	fmt.Printf("API Call for Task %d completed: %+v\n", id, post)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	// Start tracing
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Start multiple goroutines
	for i := 1; i <= 4; i++ {
		go task(i) // Start task in a new goroutine
		// go apiCall(i) // Start real API call in a new goroutine
	}

	// Wait for all goroutines to finish
	// time.Sleep(15 * time.Second) // Adjust this time as needed to ensure all tasks complete
	fmt.Println("All tasks are done!")
}
