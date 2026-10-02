// +build js,wasm
package main

import (
	"syscall/js"
)

func addDOMElement(this js.Value, args []js.Value) interface{} {
	// Get the global browser document object
	document := js.Global().Get("document")

	// Create a new <h1> element
	h1 := document.Call("createElement", "h1")
	h1.Set("innerHTML", "Hello from Go WebAssembly!")
	h1.Get("style").Set("color", "blue")

	// Append it to the body
	document.Get("body").Call("appendChild", h1)

	return nil
}

func main() {
	// Expose our Go function to the JavaScript window object
	js.Global().Set("triggerGoFunction", js.FuncOf(addDOMElement))

	// Wasm programs must not exit, or the Go runtime will crash in the browser.
	// We block the main thread forever using an empty channel.
	select {}
}