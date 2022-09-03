package main

import "fmt"

type Number int

// Change the receiver parameter to a poitner type.
func (n *Number) Double() {
	// Update the value at the pointer.
	*n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)
}
