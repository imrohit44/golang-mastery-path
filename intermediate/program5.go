package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Task completed"
	}()

	select {
	case res := <-ch:
		fmt.Println("Success:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Task took too long!")
	}
}