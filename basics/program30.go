package main

import "fmt"

type User struct {
	Name  string
	Admin bool
}

func main() {
	u := User{Name: "Alice", Admin: true}
	num := 255

	fmt.Printf("Default format: %v\n", u)
	fmt.Printf("Struct with field names: %+v\n", u)
	fmt.Printf("Go syntax representation: %#v\n", u)
	
	fmt.Printf("Data type: %T\n", u)
	fmt.Printf("Boolean format: %t\n", u.Admin)
	
	fmt.Printf("Base 10 integer: %d\n", num)
	fmt.Printf("Binary representation: %b\n", num)
	fmt.Printf("Hexadecimal representation: %x\n", num)
}