package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := 10
	maxWorkers := 3
	
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxWorkers)

	for i := 1; i <= tasks; i++ {
		wg.Add(1)
		
		go func(taskID int) {
			defer wg.Done()
			
			// Acquire token
			semaphore <- struct{}{}
			
			fmt.Printf("Worker executing task %d\n", taskID)
			time.Sleep(500 * time.Millisecond) // Simulate heavy work
			
			// Release token
			<-semaphore
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}