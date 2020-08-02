package main

import "fmt"

func main() {
	// to create an empty slice of non-size length
	// use the built-in make function
	// further reading:
	// https://blog.golang.org/slices-intro
	slicedAndDiced := make([]string, 3)
	fmt.Println("empt:", slicedAndDiced)

	// initalising a slice without the make built-in
	noMakeSlice := []string{}
	fmt.Println(noMakeSlice)

	noMakeSliceWithContents := []byte{0, 1, 2, 3}
	fmt.Println(noMakeSliceWithContents)

	// setting and getting from slices
	slicedAndDiced[0] = "a"
	slicedAndDiced[1] = "b"
	slicedAndDiced[2] = "c"

	fmt.Println("set:", slicedAndDiced)
	fmt.Println("get:", slicedAndDiced)

	fmt.Println("len:", len(slicedAndDiced))

	c := make([]string, len(slicedAndDiced))

	// testing out copying functionality
	// the copy method copes the 2nd argument
	// into the first
	copy(c, slicedAndDiced)
	fmt.Println("copy:", c)

	// testing out the slice operator
	// slice a full array into a slice (whole slice)
	fullSlice := slicedAndDiced[:]
	fmt.Println(fullSlice)

	// testing out other slices
	// partial slicing
	partialSliceOne := slicedAndDiced[0:3]
	fmt.Println(partialSliceOne)

	partialSliceTwo := slicedAndDiced[0:]
	fmt.Println(partialSliceTwo)

	partialSliceThree := slicedAndDiced[:2]
	fmt.Println(partialSliceThree)

	// demonstrating how slices can have varying lengths
	// this essentially creates a en empty slice with 3 elementsg
	// establishes a for loop iteration, iterates upon the throwaway i
	// creates a new array, then initates another for loop
	// which is equal the size of the incremented I size
	// then in the same array, adds the throws i and j
	// and sets in the 2d array
	twoDimensionalDynamicSizeSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoDimensionalDynamicSizeSlice[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDimensionalDynamicSizeSlice[i][j] = i + j
		}
		fmt.Println("varying size:", twoDimensionalDynamicSizeSlice)
	}
}
