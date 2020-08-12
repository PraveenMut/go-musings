package main

import "fmt"

func sums(nums ...int) {
	fmt.Print("Args provided: ", nums)
	total := 0
	for _, number := range nums {
		total += number
	}
	fmt.Printf("\nthe total sum is %d\n", total)
}

func main() {
	sums(1, 2)
	sums(1, 2, 3, 4)
	sums(1)
}
