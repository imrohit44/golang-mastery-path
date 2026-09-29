package main

import (
	"encoding/json"
	"fmt"
)

type UserProfile struct {
	Username string `json:"user_name"`                  // Renamed in JSON
	Email    string `json:"email,omitempty"`            // Omitted if empty ("")
	Password string `json:"-"`                          // Completely ignored
	Age      int    `json:"age,string"`                 // Converted to string in JSON
}

func main() {
	user := UserProfile{
		Username: "gopher99",
		Email:    "", // Will be omitted
		Password: "supersecretpassword",
		Age:      28,
	}

	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized JSON:")
	fmt.Println(string(jsonData))
}