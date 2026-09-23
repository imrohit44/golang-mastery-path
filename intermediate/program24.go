package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Prepare the command (e.g., 'echo Hello from the OS')
	cmd := exec.Command("echo", "Hello from the OS")

	// Run the command and capture its combined stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	fmt.Printf("Command Output: %s", string(output))
}