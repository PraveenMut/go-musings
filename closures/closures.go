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

func adder() func(int) int {
	i := 0
	return func(x int) int {
		i += x
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

	positiveSummation, negativeSummation := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println("sum is:", positiveSummation(i))
		fmt.Println("sum is:", negativeSummation(-2*i))
	}
}
