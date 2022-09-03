package main

import (
	"fmt"
)

// severalInts(values ... variable types) works as inputting values with holding a slice with the arguments.
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	severalInts(1)
	severalInts(1, 2, 3)
	severalInts(1, 2, 3, 4)
	mix(1, true, "a", "b")
}
