package main

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2026-09-26 15:30:00"
	
	// The layout string must match the format of the dateString, 
	// using the specific numbers from Go's reference date.
	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	fmt.Println("Parsed Time Object:", parsedTime)
	fmt.Println("Year:", parsedTime.Year())
	fmt.Println("Month:", parsedTime.Month())
}