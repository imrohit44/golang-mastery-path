package main

import (
	"fmt"
	"plugin"
)

func main() {
	// 1. Load the shared library
	p, err := plugin.Open("greeter.so")
	if err != nil {
		fmt.Println("Error loading plugin (did you compile it?):", err)
		return
	}

	// 2. Look up an exported symbol (function or variable)
	symGreeter, err := p.Lookup("Greet")
	if err != nil {
		panic(err)
	}

	// 3. Assert the symbol to the correct function signature
	greetFunc, ok := symGreeter.(func(string) string)
	if !ok {
		panic("Plugin symbol 'Greet' does not match expected signature")
	}

	// 4. Execute the dynamically loaded code
	result := greetFunc("Developer")
	fmt.Println("Plugin returned:", result)
}