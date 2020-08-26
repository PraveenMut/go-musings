package main

import (
	"fmt"
)

// this is an example of using recursion using golang
// this is a simple factorial
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// find fibs
func fibs(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibs(n-1) + fibs(n-2)
}

func main() {
	// and here we call the function and assign them to vars
	factorial1 := factorial(1)
	factorial2 := factorial(2)
	factorialbig := factorial(10)
	// and print them here
	fmt.Printf("the factorial of 1 is %d\n", factorial1)
	fmt.Printf("the factorial of 2 is %d\n", factorial2)
	fmt.Printf("the factorial of 10 is %d\n", factorialbig)

	// find fibs
	fmt.Println(fibs(4))
}
