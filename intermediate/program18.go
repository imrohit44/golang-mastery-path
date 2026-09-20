package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t.Format("15:04:05.000"))
		case <-done:
			ticker.Stop()
			fmt.Println("Ticker stopped")
			return
		}
	}
}