package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Bob", Age: 30}
    fmt.Printf("User %s is %d years old.\n", p.Name, p.Age)
}