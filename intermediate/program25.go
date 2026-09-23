package main

import (
	"fmt"
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC RECOVERED: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

func riskyHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("crash") == "true" {
		panic("Simulated database connection lost!")
	}
	fmt.Fprintln(w, "All systems operational. Add ?crash=true to simulate a panic.")
}

func main() {
	http.HandleFunc("/", RecoveryMiddleware(riskyHandler))
	fmt.Println("Server running on :8080")
	// http.ListenAndServe(":8080", nil)
}