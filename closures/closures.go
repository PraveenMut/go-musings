package main

import "fmt"

// functions in go can form closures
// where inline anonymous functions or lambdas are declared
// and then close over its embodying function, creating a variable
// that is bound by the closed upon functions
func closuresBaby() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// this another example of closures
// where the adder func returns another func
// whos args are of an int and return an int
func adder() func(int) int {
	i := 0
	return func(x int) int {
		i += x
		return i
	}
}

func main() {
	// creating a function and the assignining it a var
	testIterator := closuresBaby()

	// as the testIterator now embodies a function
	// it is now in round braces ()
	fmt.Println(testIterator())
	fmt.Println(testIterator())
	fmt.Println(testIterator())
	fmt.Println(testIterator())

	// this is to illustrate that the state of the closed
	// vars are purely when run. This creates a "private"
	// namespace almost. Similar to private namespaces
	// in java
	newState := closuresBaby()
	fmt.Println(newState())
	fmt.Println(newState())

	// also illustrates how you can have the exact same function
	// with different args, but yield different outputs as it
	// handles state independently
	positiveSummation, negativeSummation := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println("sum is:", positiveSummation(i))
		fmt.Println("sum is:", negativeSummation(-2*i))
	}
}
