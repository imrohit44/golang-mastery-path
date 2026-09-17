package main

import (
	"fmt"
	"sync"
)

func generateNumbers(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexed := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for val := range c {
			multiplexed <- val
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}

func main() {
	ch1 := generateNumbers(1, 3, 5)
	ch2 := generateNumbers(2, 4, 6)

	merged := fanIn(ch1, ch2)
	for val := range merged {
		fmt.Println("Merged Output:", val)
	}
}