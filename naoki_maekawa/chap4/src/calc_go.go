package main

import (
	"calc" // import packages from /Users/####/.goenv/versions/1.18.3/src/ directory
	"fmt"
)

func main() {
	fmt.Println(calc.Add(1, 2)) // call package qualifiers.
	fmt.Println(calc.Subtract(7, 3))
}
