package main

import "fmt"

func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Error: Division by zero"
    }
    return a / b, ""
}

func main() {
    result, err := divide(10, 2)
    if err != "" {
        fmt.Println(err)
    } else {
        fmt.Println("Result:", result)
    }
}