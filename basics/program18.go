package main

import "fmt"

type Circle struct {
	Radius float64
}

// Method with a receiver of type Circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func main() {
	c := Circle{Radius: 5}
	fmt.Printf("Area of the circle: %.2f\n", c.Area())
}