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

	// to iterate upon key-value pairs in maps is also
	// quite straight forward and can be done through
	// ranges specifying the data struture like with
	// an ordinary array
	// the first iterator is the key and the second
	// is te value
	mapItUpRealGood := map[string]string{"foo": "bar", "bar": "baz"}
	for k, v := range mapItUpRealGood {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// indeed, the value can be omitted and it will just
	// iterate on the keys
	for k := range mapItUpRealGood {
		fmt.Printf("the key is: %s\n", k)
	}
	// strings can be also be used too for ranges
	for i, r := range "golangbaby" {
		fmt.Println("the byte is", i, r)
	}
}
