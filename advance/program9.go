package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, burst),
	}

	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate))
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
				// Bucket is full
			}
		}
	}()

	return rl
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

func main() {
	limiter := NewRateLimiter(2, 3) // 2 ops/sec, burst size of 3

	for i := 1; i <= 5; i++ {
		limiter.Wait()
		fmt.Printf("[%s] Processed Request %d\n", time.Now().Format("15:04:05.000"), i)
	}
}