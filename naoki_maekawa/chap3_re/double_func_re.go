package main

import (
	"fmt"
)

func double(number *int) { // receive the address of variables with the pointer type
	*number *= 2 // no returnable answer, but returns the address of variables
}

func main() {
	amount := 6
	double(&amount)     // send the address of amount and receive the answer
	fmt.Println(amount) // the variable "amount" has already changed in function "double."
}
