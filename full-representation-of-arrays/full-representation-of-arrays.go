package main

import "fmt"

func main() {
	//  this initialises an empty array of 5 integers.
	//  by default, null values are zeroed by 0 ints
	var arr [5]int
	fmt.Println("emp:", arr)

	// to set a value in an array, you need to specify the index
	// through an assignment.

	arr[4] = 4
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	// calculate the length of the array
	fmt.Println("len", len(arr))

	// initialise a new array type with predef vals
	var ary = [3]int{1, 2, 3}
	fmt.Println("new ary print:", ary)

}
