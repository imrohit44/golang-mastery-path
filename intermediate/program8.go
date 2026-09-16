package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status": "healthy"}`)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server running on http://localhost:8080")
}