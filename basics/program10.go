package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printMessage("Async task running...") // Runs concurrently
    printMessage("Main execution")
}