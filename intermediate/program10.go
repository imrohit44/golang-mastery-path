package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"

	err := os.WriteFile(filename, []byte("Hello Go\nLine 2"), 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	defer os.Remove(filename) 

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Read Line:", scanner.Text())
	}
}