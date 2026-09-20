package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	
	fmt.Printf("Task %d starting\n", id)
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	fmt.Printf("Task %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter
		go processTask(i, &wg)
	}

	wg.Wait() // Block until the counter reaches zero
	fmt.Println("All tasks completed.")
}