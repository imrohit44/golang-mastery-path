package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting application...")

	// Schedule a function to run after 2 seconds
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("--> Timer fired! Executing delayed task.")
	})

	// Simulate some immediate work
	time.Sleep(1 * time.Second)
	fmt.Println("Application doing immediate work...")

	// Decide if we want to cancel the timer
	shouldCancel := false 
	if shouldCancel {
		timer.Stop()
		fmt.Println("Timer cancelled before it could fire.")
	}

	// Wait enough time for the timer to fire if not cancelled
	time.Sleep(2 * time.Second)
	fmt.Println("Application exiting.")
}