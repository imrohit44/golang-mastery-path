package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// Simulate a massive JSON array stream
	massiveJSONStream := `[
		{"id": 1, "name": "Alice"},
		{"id": 2, "name": "Bob"},
		{"id": 3, "name": "Charlie"}
	]`

	decoder := json.NewDecoder(strings.NewReader(massiveJSONStream))

	// Read the opening bracket '['
	t, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Started parsing array: %v\n", t)

	// Iterate over array elements while there are more items
	for decoder.More() {
		var user struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		
		// Decode only the current object
		if err := decoder.Decode(&user); err != nil {
			panic(err)
		}
		fmt.Printf("Streamed User -> ID: %d, Name: %s\n", user.ID, user.Name)
	}

	// Read the closing bracket ']'
	t, err = decoder.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Finished parsing array: %v\n", t)
}