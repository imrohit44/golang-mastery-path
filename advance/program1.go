package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: received cancellation signal, exiting\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: channel closed, exiting\n", id)
				return
			}
			fmt.Printf("Worker %d: executing job %d\n", id, job.ID)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	jobs := make(chan Job, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All workers shut down gracefully.")
}