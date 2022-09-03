package main

import (
	"fmt"
)

func sayHi() { // declare the function
	fmt.Println("Hi!")
}

/*
naming rule: do not declare the number for the first letter
			 do not start with the lowercase letter because it is the functions not packages
*/

func main() {
	sayHi() // call the function you have already defined.
}
