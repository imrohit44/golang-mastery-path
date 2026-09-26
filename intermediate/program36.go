package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Define a struct that holds state
type VisitCounter struct {
	mu    sync.Mutex
	count int
}

// Implement the ServeHTTP method required by the http.Handler interface
func (vc *VisitCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vc.mu.Lock()
	vc.count++
	currentCount := vc.count
	vc.mu.Unlock()

	fmt.Fprintf(w, "You are visitor number %d", currentCount)
}

func main() {
	counter := &VisitCounter{}
	
	// Register the struct as a handler
	http.Handle("/visitors", counter)
	
	fmt.Println("Server running on :8080/visitors")
	// http.ListenAndServe(":8080", nil)
}