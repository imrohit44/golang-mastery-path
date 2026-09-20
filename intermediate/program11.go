package main

import "fmt"

type Vehicle struct {
	Brand string
	Year  int
}

func (v Vehicle) Start() {
	fmt.Println(v.Brand, "engine starting...")
}

type Car struct {
	Vehicle      // Embedded struct
	NumDoors int
}

func main() {
	myCar := Car{
		Vehicle:  Vehicle{Brand: "Toyota", Year: 2022},
		NumDoors: 4,
	}

	// Access promoted fields and methods directly
	fmt.Println("Car Brand:", myCar.Brand)
	myCar.Start()
}