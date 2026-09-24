package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedResource = false

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	// Start 3 workers waiting for the signal
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cond.L.Lock()
			for !sharedResource { // Always check condition in a loop
				cond.Wait() // Pauses goroutine and unlocks mutex until signaled
			}
			fmt.Printf("Worker %d starting work!\n", id)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Broadcasting signal to wake all workers...")
	
	cond.L.Lock()
	sharedResource = true
	cond.L.Unlock()
	
	cond.Broadcast() // Wake up ALL waiting goroutines

	wg.Wait()
}