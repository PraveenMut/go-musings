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
	circle1 := circle{radius: 10}
	fmt.Println("Area by Value: ", circle1.area())
	fmt.Println("Circumference by Value: ", circle1.circumference())

	circle2 := &circle1
	fmt.Println("Area by Pointer Ref: ", circle2.area())
	fmt.Println("Circumference by Pointer Ref: ", circle2.circumference())
}
