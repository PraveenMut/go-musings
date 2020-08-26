package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	factorial1 := factorial(1)
	factorial2 := factorial(2)
	factorialbig := factorial(10)
	fmt.Printf("the factorial of 1 is %d\n", factorial1)
	fmt.Printf("the factorial of 2 is %d\n", factorial2)
	fmt.Printf("the factorial of 10 is %d\n", factorialbig)
}
