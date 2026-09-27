package main

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	queue chan struct{}
}

// NewLeakyBucket creates a bucket with a max capacity and a steady drain rate
func NewLeakyBucket(capacity int, drainInterval time.Duration) *LeakyBucket {
	bucket := &LeakyBucket{
		queue: make(chan struct{}, capacity),
	}

	// Background worker that "drains" the bucket at a constant rate
	go func() {
		ticker := time.NewTicker(drainInterval)
		for range ticker.C {
			select {
			case <-bucket.queue:
				fmt.Println("[System] Request processed (bucket drained)")
			default:
				// Bucket is empty, do nothing
			}
		}
	}()

	return bucket
}

// AddRequest attempts to add a request to the bucket. Returns false if full.
func (b *LeakyBucket) AddRequest() bool {
	select {
	case b.queue <- struct{}{}:
		return true // Request accepted
	default:
		return false // Bucket overflowing
	}
}

func main() {
	// Bucket holds max 3 requests, drains 1 request every 500ms
	bucket := NewLeakyBucket(3, 500*time.Millisecond)

	// Simulate a massive sudden burst of 5 requests
	for i := 1; i <= 5; i++ {
		accepted := bucket.AddRequest()
		if accepted {
			fmt.Printf("Request %d accepted into queue.\n", i)
		} else {
			fmt.Printf("Request %d DROPPED (Bucket Full).\n", i)
		}
	}

	// Keep main thread alive to watch the bucket drain
	time.Sleep(2 * time.Second)
}