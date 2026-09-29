package main

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

func main() {
	// Create an in-memory pipe
	reader, writer := io.Pipe()
	var wg sync.WaitGroup

	wg.Add(1)
	// Writer Goroutine
	go func() {
		defer wg.Done()
		defer writer.Close() // Crucial to unblock the reader

		data := []string{"Streaming", "data", "through", "a", "pipe."}
		for _, word := range data {
			fmt.Fprintln(writer, word)
		}
	}()

	wg.Add(1)
	// Reader Goroutine
	go func() {
		defer wg.Done()
		
		// ReadAll will read until the writer calls Close()
		bytes, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		
		fmt.Println("Reader received:")
		fmt.Println(strings.TrimSpace(string(bytes)))
	}()

	wg.Wait()
}