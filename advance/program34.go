package main

import (
	"context"
	"fmt"
	"time"
)

// CustomContext wraps a standard context to inject hardcoded domain data
type DomainContext struct {
	context.Context
	TenantID string
}

// Override the Value method to intercept specific keys
func (c *DomainContext) Value(key interface{}) interface{} {
	if key == "tenant" {
		return c.TenantID
	}
	// Fallback to the embedded parent context
	return c.Context.Value(key)
}

func process(ctx context.Context) {
	tenant := ctx.Value("tenant")
	fmt.Println("Processing request for Tenant:", tenant)
	
	// Check if context is done (e.g., from a timeout)
	if ctx.Err() != nil {
		fmt.Println("Context error:", ctx.Err())
	}
}

func main() {
	// Create a standard context with a timeout
	parentCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Wrap it in our custom struct
	customCtx := &DomainContext{
		Context:  parentCtx,
		TenantID: "ACME-Corp-992",
	}

	process(customCtx)
}