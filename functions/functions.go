package main

import "fmt"

// functions are similar to other functions in other languages
// except that go requires explicit returns and the best practice
// is to provide the type of the argument with <arg> <type>
// functions require the type of return of the function
// thus the syntax is func func_name(<arg> <type>) <output type> {}
func plus(a int, b int) int {
	return a + b
}

// for consecutive types, you can just suffix it at the end
func plusPlusPlus(a, b, c int) int {
	return a + b + c
}

// main func where everything is executed
func main() {
	result := plus(1, 2)
	fmt.Println("result is", result)
	resultFromConsecutiveArgs := plusPlusPlus(1, 2, 3)
	fmt.Println(resultFromConsecutiveArgs)
}
