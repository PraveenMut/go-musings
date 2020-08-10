package main

import "fmt"

func exampleOfMultipleReturns(param1, param2 int) (int, int) {
	return param1, param2
}

func main() {
	// to return both values
	a, b := exampleOfMultipleReturns(1, 2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%d and %d\n", a, b)

	// to just use one var from a multiple return function, use the _ identifier
	_, c := exampleOfMultipleReturns(2, 3)
	fmt.Println(c)
}
