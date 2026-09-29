package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	// --- Stream Encoding to File ---
	file, _ := os.Create("config.json")
	cfg := Config{Host: "localhost", Port: 8080}
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(cfg)
	file.Close()

	// --- Stream Decoding from File ---
	readFile, _ := os.Open("config.json")
	defer readFile.Close()
	defer os.Remove("config.json") // Cleanup

	var loadedCfg Config
	decoder := json.NewDecoder(readFile)
	_ = decoder.Decode(&loadedCfg)

	fmt.Printf("Loaded Config: %+v\n", loadedCfg)
}