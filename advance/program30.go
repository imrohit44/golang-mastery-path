package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Register a callback to fire when the context dies
	stopFunc := context.AfterFunc(ctx, func() {
		fmt.Println("\n[SYSTEM] Context deadline exceeded. Cleaning up resources...")
	})

	fmt.Println("Doing heavy work...")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Work step %d\n", i)
	}

	// If the work finished early, you can stop the AfterFunc from firing
	stopFunc() 
}