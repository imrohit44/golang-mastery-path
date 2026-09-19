package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		fmt.Println("Weekend!")
	}
}