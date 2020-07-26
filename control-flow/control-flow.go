package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	if var1, var2 := 3, 2; var1 > 0 && var2 > 0 {
		fmt.Println("this is a positive integer")
	} else if var1%2 != 0 {
		fmt.Println(var1, "the first digit is odd")
	} else {
		fmt.Println(var1)
	}
}
