package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Secret Data: 12345"

	// Encode to Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("Encoded:", encoded)

	// Decode from Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	
	fmt.Println("Decoded:", string(decodedBytes))
}