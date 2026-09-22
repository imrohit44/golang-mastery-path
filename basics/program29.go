package main

import "fmt"

// Define custom types
type Celsius float64
type Fahrenheit float64

// Method bound to the custom type
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

func main() {
	var currentTemp Celsius = 25.0
	
	// currentTemp = 77.0 (Would cause a type mismatch error if assigned a raw float or Fahrenheit)
	
	converted := currentTemp.ToFahrenheit()
	
	fmt.Printf("%.1f°C is %.1f°F\n", currentTemp, converted)
}