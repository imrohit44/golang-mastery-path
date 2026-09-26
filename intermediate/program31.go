package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 1. Write to a CSV file
	file, err := os.Create("users.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer os.Remove("users.csv") // Cleanup for demo

	writer := csv.NewWriter(file)
	writer.Write([]string{"Name", "Role"})
	writer.Write([]string{"Alice", "Admin"})
	writer.Write([]string{"Bob", "User"})
	writer.Flush() // Ensure all data is written

	// 2. Read from the CSV file
	readFile, _ := os.Open("users.csv")
	defer readFile.Close()

	reader := csv.NewReader(readFile)
	records, _ := reader.ReadAll()

	fmt.Println("CSV Contents:")
	for _, row := range records {
		fmt.Printf("- %s is a %s\n", row[0], row[1])
	}
}