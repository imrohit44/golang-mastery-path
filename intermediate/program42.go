package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeConfig struct {
	mu    sync.RWMutex
	theme string
}

func (c *SafeConfig) GetTheme() string {
	c.mu.RLock() // Acquire read lock (allows other readers)
	defer c.mu.RUnlock()
	return c.theme
}

func (c *SafeConfig) SetTheme(newTheme string) {
	c.mu.Lock() // Acquire write lock (blocks everyone else)
	defer c.mu.Unlock()
	c.theme = newTheme
}

func main() {
	config := &SafeConfig{theme: "dark"}
	var wg sync.WaitGroup

	// Start 5 concurrent readers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Reader %d sees: %s\n", id, config.GetTheme())
		}(i)
	}

	// Start 1 concurrent writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		fmt.Println("--> Writer changing theme to light")
		config.SetTheme("light")
	}()

	wg.Wait()
}