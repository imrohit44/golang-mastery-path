package main

import (
	"context"
	"fmt"
)

// Use a custom unexported type for context keys to prevent collisions
type contextKey string
const requestIDKey contextKey = "requestID"

func processRequest(ctx context.Context) {
	// Extract the value, asserting it to the correct type
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Printf("Processing operation for Request ID: %s\n", reqID)
	} else {
		fmt.Println("No Request ID found in context")
	}
}

func main() {
	// Inject data into the context
	ctx := context.WithValue(context.Background(), requestIDKey, "REQ-9981-ABCD")
	
	processRequest(ctx)
}