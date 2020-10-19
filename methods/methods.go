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

func main() {
	circle1 := circle{radius: 10}
	fmt.Println(circle1.area())
}
