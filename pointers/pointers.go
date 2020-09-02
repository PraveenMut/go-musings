package main

import (
	"fmt"
)

func zeroByValue(ivalue int) {
	ivalue = 0
}

func zeroByPointer(ivalue *int) {
	*ivalue = 0
}

func main() {
	i := 1
	fmt.Println("initial value is:", i)

	zeroByValue(i)
	fmt.Println("the value of passing a value in is:", i)

	zeroByPointer(&i)
	fmt.Println("the value of i is now overwritten:", i)

	fmt.Println("the memory address is", &i)
}
