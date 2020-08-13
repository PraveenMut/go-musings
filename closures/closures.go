package main

import "fmt"

func closuresBaby() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	testIterator := closuresBaby()

	fmt.Println(testIterator())
	fmt.Println(testIterator())
	fmt.Println(testIterator())
	fmt.Println(testIterator())

	newState := closuresBaby()
	fmt.Println(newState())
	fmt.Println(newState())
}
