package main

import (
	"fmt"
	"math"
)

const str string = "constant"
const notStr int = 10

func main() {
	fmt.Println(str)

	const largeInt = 50000000

	const sciNotation = 3e20 / largeInt

	fmt.Println(sciNotation)
	fmt.Println(int64(sciNotation))
	fmt.Println(math.Sin(sciNotation))
}
