package main

import (
	"context"
	"fmt"
	"time"
)

func fireAndForgetLog(ctx context.Context, msg string) {
	// Simulate writing to a database that takes 2 seconds
	time.Sleep(2 * time.Second)
	
	if ctx.Err() != nil {
		fmt.Println("Failed to write log! Context was cancelled:", ctx.Err())
		return
	}
	
	traceID := ctx.Value("trace_id")
	fmt.Printf("Successfully wrote background log [%s]: %s\n", traceID, msg)
}

func main() {
	// Simulate an incoming HTTP request with a strict 1-second timeout
	reqCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	reqCtx = context.WithValue(reqCtx, "trace_id", "REQ-12345")
	defer cancel()

	// Detach the context: Keeps the "trace_id" but ignores the 1-second timeout
	detachedCtx := context.WithoutCancel(reqCtx)

	// Spawn the background task
	go fireAndForgetLog(detachedCtx, "User logged in")

	// Wait for the main request context to timeout (simulating the request ending)
	<-reqCtx.Done()
	fmt.Println("Main request finished and cancelled its context.")

	// Give the background task time to finish to prove it survived
	time.Sleep(2 * time.Second)
}