package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
    }
    
    scores["Charlie"] = 95 

    for name, score := range scores {
        fmt.Printf("%s scored %d\n", name, score)
    }
}