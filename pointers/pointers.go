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

	testValue := "share"
	testValueByPointer := &testValue

	fmt.Println("testValue is:", testValue)
	fmt.Println("testValue by Pointer address is :", testValueByPointer)
	fmt.Println("testValue by Pointer is:", *testValueByPointer)
}
