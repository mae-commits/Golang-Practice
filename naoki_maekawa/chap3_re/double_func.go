package main

import (
	"fmt"
)

func double(number int) {
	number *= 2 // no returnable things
	// return number
}

func main() {
	amount := 6
	double(amount)
	/*
		function "double" does not have returnable things, so x cannot receive values.
	*/
	fmt.Println(amount) // the variable "amount" has already changed in function "double."
	// fmt.Println(x)
}
