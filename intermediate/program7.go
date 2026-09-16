package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"user_name"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	jsonData := `{"id": 101, "user_name": "rohit_m", "tags": ["go", "dev"]}`

	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Parsed Struct: %+v\n", user)

	encoded, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("JSON Output:\n", string(encoded))
}