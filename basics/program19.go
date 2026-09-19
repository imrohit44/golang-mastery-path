package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	fmt.Println("Raw Time:", now)
	fmt.Println("Formatted Date:", now.Format("2006-01-02"))
	fmt.Println("Formatted Time:", now.Format("15:04:05"))
}