package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	/*
		Print = print with formatting, a string that will be used to format the output.
		% = a sign of starting formatting verb
		0.2 = the value widths
		f = floating-point number
		\n = escape the new sequence following \n
	*/
	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0) // only returns a formatted string.
	fmt.Printf(resultString)
	fmt.Printf("%d\n", 1)
	fmt.Printf("A string: %s\n", "String")
	fmt.Printf("%v %v %v", "", "\t", "\n")
	fmt.Printf("%#v %#v %#v\n", "", "\t", "\n")
	fmt.Printf("%f litters needed\n", 1.81999999999999)
	/*
		%#v returns the value itself.
		%v returns the string itself.
		%f returns the floating point values until 6 decimal.

	*/
}
