package main

import "fmt"

func main() {
    number := 7

    if number%2 == 0 {
        fmt.Println(number, "is Even")
    } else {
        fmt.Println(number, "is Odd")
    }
}