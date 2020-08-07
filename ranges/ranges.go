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
	// we don't want the index, so we just specify _
	for _, v := range iterableSlice {
		sum += sum + v
	}
	fmt.Println("total sum is:", sum)

	// if we want the index then we just specify it
	for i, v := range iterableSlice {
		if v%2 == 0 {
			fmt.Println("this index position is even", i)
		}
	}

	mapItUpRealGood := map[string]string{"foo": "bar", "bar": "baz"}
	for k, v := range mapItUpRealGood {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range mapItUpRealGood {
		fmt.Printf("the key is: %s\n", k)
	}
}
