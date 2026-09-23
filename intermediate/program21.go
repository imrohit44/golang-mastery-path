package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	config map[string]string
)

func loadConfig() {
	fmt.Println("--- Loading configuration from disk ---")
	config = map[string]string{"env": "production"}
}

func GetConfig() map[string]string {
	// loadConfig will only be called the very first time GetConfig runs
	once.Do(loadConfig)
	return config
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := GetConfig()
			fmt.Printf("Worker %d got env: %s\n", id, cfg["env"])
		}(i)
	}
	wg.Wait()
}