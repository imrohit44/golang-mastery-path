package main

import (
	"fmt"
	"sync"
	"time"
)

type Call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*Call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.calls == nil {
		g.calls = make(map[string]*Call)
	}
	
	// If a call is already in flight for this key, wait for it
	if c, ok := g.calls[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(Call)
	c.wg.Add(1)
	g.calls[key] = c
	g.mu.Unlock()

	// Execute the function once
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()

	return c.val, c.err
}

func expensiveDBQuery() (interface{}, error) {
	fmt.Println(">>> Executing EXPENSIVE database query...")
	time.Sleep(1 * time.Second)
	return "User Data 101", nil
}

func main() {
	var g Group
	var wg sync.WaitGroup

	// Simulate 5 identical concurrent requests
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			val, _ := g.Do("user_101", expensiveDBQuery)
			fmt.Printf("Worker %d got: %v\n", workerID, val)
		}(i)
	}

	wg.Wait()
	// Notice in the output that "Executing EXPENSIVE database query..." only prints ONCE.
}