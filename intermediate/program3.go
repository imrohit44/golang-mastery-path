package main

import "fmt"

func produce(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i * 10
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go produce(ch)

	for val := range ch {
		fmt.Println("Received:", val)
	}
}