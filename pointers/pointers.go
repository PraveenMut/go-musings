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

	*testValueByPointer = "newValue"

	fmt.Println("the new testValue is:", testValue)
	fmt.Println("the reference by pointer value is:", *testValueByPointer)

	// deciding when to pass a pointer than sending a value is about knowing
	// if you want to cause mutation in the passing parameter. If you wish
	// to mutate the parameter, then defining a function with pointer args
	// is a good approach.

	// However, if you which to scope your functions and provide scoped vars
	// then providing parameters into the function args should be done by
	// value.
}
