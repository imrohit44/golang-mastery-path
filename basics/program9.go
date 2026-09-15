package main

import "fmt"

func updateValue(val *int) {
    *val = 42 // Modify value at memory address
}

func main() {
    num := 10
    fmt.Println("Before:", num)

    updateValue(&num)
    fmt.Println("After:", num)
}