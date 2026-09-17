package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	shutdownSignal := make(chan os.Signal, 1)
	// Register system interrupts
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Service running... Press Ctrl+C to terminate.")
		for {
			time.Sleep(500 * time.Millisecond)
		}
	}()

	sig := <-shutdownSignal
	fmt.Printf("\nReceived OS Signal: %s. Starting cleanup...\n", sig)

	// Simulate resource cleanup
	time.Sleep(1 * time.Second)
	fmt.Println("Cleanup finished. Process exited cleanly.")
}