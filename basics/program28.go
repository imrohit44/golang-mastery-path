package main

import "fmt"

var appConfig string

// init always runs before main
func init() {
	fmt.Println("1. init() function is running...")
	appConfig = "Loaded Configuration"
}

func main() {
	fmt.Println("2. main() function is running...")
	fmt.Println("App Status:", appConfig)
}