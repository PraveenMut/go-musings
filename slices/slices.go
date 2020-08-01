package main

import "fmt"

func main() {
	// to create an empty slice of non-size length
	// use the built-in make
	slicedAndDiced := make([]string, 3)
	fmt.Println("empt:", slicedAndDiced)

	// setting and getting from slices
	slicedAndDiced[0] = "a"
	slicedAndDiced[1] = "b"
	slicedAndDiced[2] = "c"

	fmt.Println("set:", slicedAndDiced)
	fmt.Println("get:", slicedAndDiced)

	fmt.Println("len:", len(slicedAndDiced))

	c := make([]string, len(slicedAndDiced))

	copy(c, slicedAndDiced)
	fmt.Println("copy:", c)

	fullSlice := slicedAndDiced[:]
	fmt.Println(fullSlice)

	partialSliceOne := slicedAndDiced[0:4]
	fmt.Println(partialSliceOne)

	partialSliceTwo := slicedAndDiced[0:]
	fmt.Println(partialSliceTwo)

	partialSliceThree := slicedAndDiced[:2]
	fmt.Println(partialSliceThree)

}
