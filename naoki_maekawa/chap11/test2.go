package main

import (
	"fmt"
	"mypkg"
)

func main() {
	// Declare a variable using the interface type.
	var value mypkg.MyInterface
	// Values of myType satisfy myInterface, so we can assign this value to a variable with a type of myInterface.
	value = mypkg.MyType(5)
	// We can call any method that is part of myInterface.
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
