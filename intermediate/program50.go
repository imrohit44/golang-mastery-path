package main

import "fmt"

type DefaultLogger struct{}

func (l DefaultLogger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

// SecureLogger embeds DefaultLogger
type SecureLogger struct {
	DefaultLogger
}

// Override the Log method for SecureLogger
func (sl SecureLogger) Log(msg string) {
	fmt.Println("[SECURE-LOG] ***REDACTED***")
}

func main() {
	basic := DefaultLogger{}
	basic.Log("User logged in") // Calls DefaultLogger's method

	secure := SecureLogger{}
	secure.Log("User logged in") // Calls overridden SecureLogger method
	
	// We can still explicitly call the underlying embedded method
	secure.DefaultLogger.Log("Admin bypass active") 
}