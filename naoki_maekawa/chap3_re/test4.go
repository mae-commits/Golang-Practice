package main

import (
	"fmt"
)

func sayHi() { // declare the function
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) { // declare the function with variables (= parameters)
	// line, times = variables' name, string,int = types of each varible for line and times, respectively
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

/*
naming rule: do not declare the number for the first letter
			 do not start with the lowercase letter because it is the functions not packages
*/
func main() {
	sayHi() // call the function you have already defined.
	repeatLine("hello", 3)
}
