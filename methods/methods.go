package main

import (
	"fmt"
	"math"
)

// create a basic circle struct to allow for calculation of area and perimeter
type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// Instantiate a struct from circle
	circle1 := circle{radius: 10}
	// Implement a method call on the newly instantiated circle1 struct by value.

	fmt.Println("Area by Value: ", circle1.area())
	// This will not mutate the incoming struct as it is passed by Value
	fmt.Println("Circumference by Value: ", circle1.circumference())

	// Impelement a method but in this case pass by pointer
	circle2 := &circle1

	// this has the ability to mutate the receiving struct as it passed by pointer.
	fmt.Println("Area by Pointer Ref: ", circle2.area())
	fmt.Println("Circumference by Pointer Ref: ", circle2.circumference())
}
