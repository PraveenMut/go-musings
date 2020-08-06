package main

import (
	"fmt"
)

func main() {
	// range allow you to iterate over data structures
	// similar to the `in` python or `.each` in ruby or `.foreach(e =>)` in js

	// take a slice iterableSlice
	iterableSlice := []int{1, 2, 3}
	// initalise a sum var
	sum := 0
	// you can use a for loop to start an iterable process
	// the for loop takes 2 enumerators/iterators, first
	// is the index and the second is the value itself
	for _, v := range iterableSlice {
		sum += sum + v
	}
	fmt.Println("total sum is:", sum)
}
