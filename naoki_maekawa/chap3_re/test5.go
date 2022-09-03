package main

import (
	"fmt"
)

func double(number float64) float64 {
	return number * 2 // return keyboard, and exits immediately, so the following code are not read
}

func main() {
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}
