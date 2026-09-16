package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (c Circle) Area() float64    { return 3.14159 * c.Radius * c.Radius }

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 4}

	printArea(rect)
	printArea(circ)
}